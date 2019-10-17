package container

import (
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/sportmonks"
)

func (c Container) CountryRequester() app.CountryRequester {
	return sportmonks.NewCountryRequester(c.NewSportMonksClient, c.NewLogger)
}
