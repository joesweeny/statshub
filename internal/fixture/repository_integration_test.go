package fixture

import (
	"database/sql"
	"fmt"
	"github.com/joesweeny/statshub/internal/config"
	"github.com/joesweeny/statshub/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresFixtureRepository{Connection: conn}

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		for i := 1; i < 4; i++ {
			c := newFixture(i)

			if err := repo.Insert(c); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}

			row := conn.QueryRow("select count(*) from sportmonks_fixture")

			var count int

			if err := row.Scan(&count); err != nil {
				t.Errorf("Error when scanning rows returned by the database: %s", err.Error())
			}

			assert.Equal(t, i, count)
		}
	})

	t.Run("returns error when ID primary key violates unique constraint", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newFixture(50)

		if err := repo.Insert(c); err != nil {
			t.Errorf("Test failed, expected nil, got %s", err)
		}

		if e := repo.Insert(c); e == nil {
			t.Fatalf("Test failed, expected %s, got nil", e)
		}
	})

	conn.Close()
}

func TestGetById(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresFixtureRepository{Connection: conn}

	t.Run("fixture can be retrieved by ID", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		f := newFixture(43)

		if err := repo.Insert(f); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		r, err := repo.GetById(43)

		if err != nil {
			t.Errorf("Error when retrieving a record from the database: %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(43, r.ID)
		a.Equal(14567, r.SeasonID)
		a.Equal(165789, *r.RoundID)
		a.Nil(f.VenueID)
		a.Equal(451, r.HomeTeamID)
		a.Equal(924, r.AwayTeamID)
		a.Nil(r.RefereeID)
		a.Equal("2019-01-21 16:08:49 +0000 UTC", r.Date.String())
		a.Equal("2019-01-08 16:33:20 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-08 16:33:20 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns error if fixture does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		_, err := repo.GetById(99)

		if err == nil {
			t.Errorf("Test failed, expected %v, got nil", err)
		}

		if err != ErrNotFound {
			t.Fatalf("Test failed, expected %v, got %s", ErrNotFound, err)
		}
	})

	conn.Close()
}

func TestUpdate(t *testing.T) {
	conn, cleanUp := getConnection(t)
	repo := PostgresFixtureRepository{Connection: conn}

	t.Run("modifies existing fixture", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		f := newFixture(78)

		if err := repo.Insert(f); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		var venueId = 574
		var roundId *int
		var d = time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC)

		f.VenueID = &venueId
		f.AwayTeamID = 4390
		f.RoundID = roundId
		f.Date = d

		if err := repo.Update(f); err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		r, err := repo.GetById(78)

		if err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(78, f.ID)
		a.Equal(14567, f.SeasonID)
		a.Nil(f.RoundID)
		a.Equal(574, *f.VenueID)
		a.Equal(451, f.HomeTeamID)
		a.Equal(4390, f.AwayTeamID)
		a.Nil(f.RefereeID)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.Date.String())
		a.Equal("2019-01-08 16:33:20 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-08 16:33:20 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns an error if fixture does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newFixture(146)

		err := repo.Update(c)

		if err == nil {
			t.Fatalf("Test failed, expected nil, got %v", err)
		}

		if err != ErrNotFound {
			t.Fatalf("Test failed, expected %v, got %v", ErrNotFound, err)
		}
	})

	conn.Close()
}

var db = config.GetConfig().Database

func getConnection(t *testing.T) (*sql.DB, func()) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		db.Host, db.Port, db.User, db.Password, db.Name)

	db, err := sql.Open(db.Driver, psqlInfo)

	if err != nil {
		panic(err)
	}

	return db, func() {
		_, err := db.Exec("delete from sportmonks_fixture")
		if err != nil {
			t.Fatalf("Failed to clear database. %s", err.Error())
		}
	}
}

func newFixture(id int) *model.Fixture {
	var roundId = 165789

	return &model.Fixture{
		ID:         id,
		SeasonID:   14567,
		RoundID:    &roundId,
		HomeTeamID: 451,
		AwayTeamID: 924,
		Date:       time.Unix(1548086929, 0),
		CreatedAt:  time.Unix(1546965200, 0),
		UpdatedAt:  time.Unix(1546965200, 0),
	}
}