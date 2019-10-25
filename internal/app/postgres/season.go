package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jonboulle/clockwork"
	"github.com/statistico/statistico-data/internal/app"
	"time"
)

type SeasonRepository struct {
	connection *sql.DB
	clock      clockwork.Clock
}

func (r *SeasonRepository) Insert(s *app.Season) error {
	query := `
	INSERT INTO sportmonks_season (id, name, league_id, is_current, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.connection.Exec(
		query,
		s.ID,
		s.Name,
		s.CompetitionID,
		s.IsCurrent,
		s.CreatedAt.Unix(),
		s.UpdatedAt.Unix(),
	)

	return err
}

func (r *SeasonRepository) Update(s *app.Season) error {
	_, err := r.ByID(s.ID)

	if err != nil {
		return err
	}

	query := `
	 UPDATE sportmonks_season set name = $2, league_id = $3, is_current = $4, updated_at = $5 where id = $1`

	_, err = r.connection.Exec(
		query,
		s.ID,
		s.Name,
		s.CompetitionID,
		s.IsCurrent,
		s.UpdatedAt.Unix(),
	)

	return err
}

func (r *SeasonRepository) ByID(id int64) (*app.Season, error) {
	query := `SELECT * FROM sportmonks_season where id = $1`
	row := r.connection.QueryRow(query, id)

	return rowToSeason(row)
}

func (r *SeasonRepository) IDs() ([]int64, error) {
	query := `SELECT id FROM sportmonks_season ORDER BY id ASC`

	rows, err := r.connection.Query(query)

	if err != nil {
		return []int64{}, err
	}

	defer rows.Close()

	var id int64
	var ids []int64

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func (r *SeasonRepository) CurrentSeasonIDs() ([]int64, error) {
	query := `SELECT id FROM sportmonks_season where is_current = true ORDER BY id ASC`

	rows, err := r.connection.Query(query)

	if err != nil {
		return []int64{}, err
	}

	defer rows.Close()

	var seasons []int64

	for rows.Next() {
		var id int64

		if err := rows.Scan(&id); err != nil {
			return seasons, err
		}

		seasons = append(seasons, id)
	}

	return seasons, nil
}

func rowToSeason(r *sql.Row) (*app.Season, error) {
	var created int64
	var updated int64

	var s = app.Season{}

	if err := r.Scan(&s.ID, &s.Name, &s.CompetitionID, &s.IsCurrent, &created, &updated); err != nil {
		return &s, errors.New(fmt.Sprintf("Season with ID %d does not exist", s.ID))
	}

	s.CreatedAt = time.Unix(created, 0)
	s.UpdatedAt = time.Unix(updated, 0)

	return &s, nil
}
