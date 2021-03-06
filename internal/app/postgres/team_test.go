package postgres_test

import (
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/errors"
	"github.com/statistico/statistico-data/internal/app/postgres"
	"github.com/statistico/statistico-data/internal/app/test"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTeamRepository_Insert(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_team")
	repo := postgres.NewTeamRepository(conn, test.Clock)

	t.Run("increases table count", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		for i := 1; i < 4; i++ {
			c := newTeam(uint64(i), "West Ham United")

			if err := repo.Insert(c); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}

			row := conn.QueryRow("select count(*) from sportmonks_team")

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
		c := newTeam(50, "West Ham United")

		if err := repo.Insert(c); err != nil {
			t.Errorf("Test failed, expected nil, got %s", err)
		}

		if e := repo.Insert(c); e == nil {
			t.Fatalf("Test failed, expected %s, got nil", e)
		}
	})
}

func TestTeamRepository_ByID(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_team")
	repo := postgres.NewTeamRepository(conn, test.Clock)

	t.Run("team can be retrieved by ID", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		m := newTeam(43, "West Ham United")

		if err := repo.Insert(m); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		r, err := repo.ByID(43)

		if err != nil {
			t.Errorf("Error when retrieving a record from the database: %s", err.Error())
		}

		if r == nil {
			t.Fatalf("Expected resource, got nil")
		}

		a := assert.New(t)

		a.Equal(uint64(43), r.ID)
		a.Equal("West Ham United", r.Name)
		a.Equal(uint64(560), r.VenueID)
		a.Equal(false, r.NationalTeam)
		a.Nil(r.ShortCode)
		a.Equal(uint64(462), r.CountryID)
		a.Nil(r.Founded)
		a.Nil(r.Logo)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns error if team does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		_, err := repo.ByID(4)

		if err == nil {
			t.Fatalf("Test failed, expected %v, got nil", err)
		}

		assert.Equal(t, errors.ErrorNotFound, err)
	})
}

func TestTeamRepository_Update(t *testing.T) {
	conn, cleanUp := test.GetConnection(t, "sportmonks_team")
	repo := postgres.NewTeamRepository(conn, test.Clock)

	t.Run("modifies existing team", func(t *testing.T) {
		t.Helper()
		defer cleanUp()

		m := newTeam(43, "West Ham United")

		if err := repo.Insert(m); err != nil {
			t.Errorf("Error when inserting record into the database: %s", err.Error())
		}

		var shortCode = "WHU"
		var founded = 1898
		var logo = "http://path.to/logo"
		var d = time.Date(2019, 01, 14, 11, 25, 00, 00, time.UTC)

		m.Name = "West Ham London Boooo"
		m.ShortCode = &shortCode
		m.CountryID = uint64(5)
		m.Founded = &founded
		m.Logo = &logo
		m.UpdatedAt = d

		if err := repo.Update(m); err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		r, err := repo.ByID(43)

		if err != nil {
			t.Errorf("Error when updating a record in the database: %s", err.Error())
		}

		if r == nil {
			t.Fatalf("Expected resource, got nil")
		}

		a := assert.New(t)

		a.Equal(uint64(43), r.ID)
		a.Equal("West Ham London Boooo", r.Name)
		a.Equal(uint64(560), r.VenueID)
		a.Equal(false, r.NationalTeam)
		a.Equal("WHU", *r.ShortCode)
		a.Equal(uint64(5), r.CountryID)
		a.Equal(1898, *r.Founded)
		a.Equal("http://path.to/logo", *r.Logo)
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.CreatedAt.String())
		a.Equal("2019-01-14 11:25:00 +0000 UTC", r.UpdatedAt.String())
	})

	t.Run("returns an error if team does not exist", func(t *testing.T) {
		t.Helper()
		defer cleanUp()
		c := newTeam(146, "West Ham United")

		err := repo.Update(c)

		if err == nil {
			t.Fatalf("Test failed, expected %v, got nil", err)
		}

		assert.Equal(t, errors.ErrorNotFound, err)
	})
}

func TestTeamRepository_BySeasonId(t *testing.T) {
	teamConn, teamCleanUp := test.GetConnection(t, "sportmonks_team")
	fixtureConn, fixtureCleanUp := test.GetConnection(t, "sportmonks_fixture")
	teamRepo := postgres.NewTeamRepository(teamConn, test.Clock)
	fixtureRepo := postgres.NewFixtureRepository(fixtureConn, test.Clock)

	t.Run("returns a slice on team struct of teams associated to a season", func(t *testing.T) {
		t.Helper()
		defer teamCleanUp()
		defer fixtureCleanUp()

		teams := []*app.Team{
			newTeam(1, "Newcastle United"),
			newTeam(3, "West Ham United"),
			newTeam(2, "AFC Bournemouth"),
			newTeam(4, "Chelsea"),
		}

		for _, team := range teams {
			err := teamRepo.Insert(team)

			if err != nil {
				t.Fatalf("Expected nil, got %s", err.Error())
			}
		}

		fixtures := []*app.Fixture{
			{
				ID:         1,
				SeasonID:   uint64(14567),
				HomeTeamID: 1,
				AwayTeamID: 3,
				Date:       time.Unix(1548086929, 0),
			},
			{
				ID:         2,
				SeasonID:   uint64(14567),
				HomeTeamID: 1,
				AwayTeamID: 4,
				Date:       time.Unix(1548086920, 0),
			},
			{
				ID:         3,
				SeasonID:   uint64(14567),
				HomeTeamID: 3,
				AwayTeamID: 1,
				Date:       time.Unix(1548086925, 0),
			},
			{
				ID:         4,
				SeasonID:   uint64(14567),
				HomeTeamID: 4,
				AwayTeamID: 3,
				Date:       time.Unix(1548086925, 0),
			},
			{
				ID:         5,
				SeasonID:   uint64(65432),
				HomeTeamID: 2,
				AwayTeamID: 4,
				Date:       time.Unix(1548086925, 0),
			},
		}

		for _, fix := range fixtures {
			if err := fixtureRepo.Insert(fix); err != nil {
				t.Errorf("Error when inserting record into the database: %s", err.Error())
			}
		}

		fetched, err := teamRepo.BySeasonId(14567)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, 3, len(fetched))
		assert.Equal(t, uint64(4), fetched[0].ID)
		assert.Equal(t, "Chelsea", fetched[0].Name)
		assert.Equal(t, uint64(1), fetched[1].ID)
		assert.Equal(t, "Newcastle United", fetched[1].Name)
		assert.Equal(t, uint64(3), fetched[2].ID)
		assert.Equal(t, "West Ham United", fetched[2].Name)
	})
}

func newTeam(id uint64, name string) *app.Team {
	return &app.Team{
		ID:           id,
		Name:         name,
		VenueID:      560,
		CountryID:    uint64(462),
		NationalTeam: false,
		CreatedAt:    time.Unix(1546965200, 0),
		UpdatedAt:    time.Unix(1546965200, 0),
	}
}
