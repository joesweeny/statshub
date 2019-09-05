package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/statistico/statistico-data/internal/statistico"
	"time"
)

type CountryRepository struct {
	Connection *sql.DB
}

// Insert a new domain Country struct to database
// Errors that occur while performing the operation are returned.
func (p *CountryRepository) Insert(c *statistico.Country) error {
	query := `
	INSERT INTO sportmonks_country (id, name, continent, iso, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := p.Connection.Exec(
		query,
		c.ID,
		c.Name,
		c.Continent,
		c.ISO,
		c.CreatedAt.Unix(),
		c.UpdatedAt.Unix(),
	)

	return err
}

// Update an existing domain Country struct to database
// Errors that occur while performing the operation are returned.
func (p *CountryRepository) Update(c *statistico.Country) error {
	_, err := p.GetById(c.ID)

	if err != nil {
		return err
	}

	query := `
	UPDATE sportmonks_country 
	set name = $2, continent = $3, iso = $4, updated_at = $5 
	where id = $1`

	_, err = p.Connection.Exec(
		query,
		c.ID,
		c.Name,
		c.Continent,
		c.ISO,
		c.UpdatedAt.Unix(),
	)

	return err
}

// Retrieve an existing domain Country struct from database
// Errors that occur while performing the operation are returned.
func (p *CountryRepository) GetById(id int) (*statistico.Country, error) {
	query := `SELECT * from sportmonks_country where id = $1`
	row := p.Connection.QueryRow(query, id)

	return rowToCountry(row)
}

func rowToCountry(r *sql.Row) (*statistico.Country, error) {
	var created int64
	var updated int64

	c := statistico.Country{}

	if err := r.Scan(&c.ID, &c.Name, &c.Continent, &c.ISO, &created, &updated); err != nil {
		return &c, errors.New(fmt.Sprintf("Country with ID %d does not exist", c.ID))
	}

	c.CreatedAt = time.Unix(created, 0)
	c.UpdatedAt = time.Unix(updated, 0)

	return &c, nil
}
