package rest

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/statistico/statistico-data/internal/app"
	"net/http"
	"net/url"
	"strconv"
	"time"
)
type FixtureHandler struct {
	fixtureRepo app.FixtureRepository
	factory FixtureFactory
}

func (f FixtureHandler) seasonFixtures(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	query, err := parseFixtureQuery(r, ps)

	if err == errBadRequest {
		failResponse(w, http.StatusBadRequest, err)
	}

	if err == errTimeParse {
		failResponse(w, http.StatusUnprocessableEntity, err)
	}

	fixtures, err := f.fixtureRepo.Get(query)

	if err != nil {
		errorResponse(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	var response []*Fixture

	for _, fix := range fixtures {
		f, err := f.factory.BuildFixture(&fix)

		if err != nil {
			errorResponse(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
			return
		}

		response = append(response, f)
	}

	successResponse(w, http.StatusOK, response)
}

func parseFixtureQuery(r *http.Request, ps httprouter.Params) (app.FixtureRepositoryQuery, error) {
	query := app.FixtureRepositoryQuery{}

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		return query, errBadRequest
	}

	after, err := parseDateQuery(r.URL.Query(), "date_after")

	if err == errTimeParse {
		return query, err
	}

	before, err := parseDateQuery(r.URL.Query(), "date_before")

	if err == errTimeParse {
		return query, err
	}

	seasonID := uint64(id)
	query.SeasonID =  &seasonID
	query.DateFrom =   after
	query.DateTo =     before

	return query, nil
}

func parseDateQuery(query url.Values, key string) (*time.Time, error) {
	val := query.Get(key)

	if val == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, val)

	if err != nil {
		return nil, errTimeParse
	}

	return &t, nil
}

