package grpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/grpc/factory"
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type ResultService struct {
	fixtureRepo app.FixtureRepository
	factory     *factory.ResultFactory
	logger      *logrus.Logger
	statistico.UnimplementedResultServiceServer
}

func (s ResultService) GetById(ctx context.Context, r *statistico.ResultRequest) (*statistico.Result, error) {
	fix, err := s.fixtureRepo.ByID(r.GetFixtureId())

	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Result for fixture %d does not exist", r.GetFixtureId()))
	}

	x, err := s.factory.BuildResult(fix)

	if err != nil {
		s.logger.Warnf("Error hydrating Result. Error: %s", err.Error())
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return x, nil
}

func (s ResultService) GetHistoricalResultsForFixture(r *statistico.HistoricalResultRequest, stream statistico.ResultService_GetHistoricalResultsForFixtureServer) error {
	date, err := time.Parse(time.RFC3339, r.DateBefore)

	if err != nil {
		return status.Error(codes.InvalidArgument, "Date provided is not a valid RFC3339 date")
	}

	limit := uint64(r.Limit)

	query := app.FixtureRepositoryQuery{
		HomeTeamID: &r.HomeTeamId,
		AwayTeamID: &r.AwayTeamId,
		DateTo:     &date,
		Limit:      &limit,
	}

	fixtures, err := s.fixtureRepo.Get(query)

	if err != nil {
		s.logger.Warnf("Error retrieving Fixture(s) in Result Service. Error: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	return s.sendResults(fixtures, stream)
}

func (s ResultService) GetResultsForTeam(r *statistico.TeamResultRequest, stream statistico.ResultService_GetResultsForTeamServer) error {
	var query app.FixtureFilterQuery

	if r.GetDateBefore() != nil {
		date, err := time.Parse(time.RFC3339, r.GetDateBefore().GetValue())

		if err != nil {
			return status.Error(
				codes.InvalidArgument,
				fmt.Sprintf("Date provided '%s' is not a valid RFC3339 date", r.GetDateBefore().GetValue()),
			)
		}

		query.DateBefore = &date
	}

	if r.GetDateAfter() != nil {
		date, err := time.Parse(time.RFC3339, r.GetDateAfter().GetValue())

		if err != nil {
			return status.Error(
				codes.InvalidArgument,
				fmt.Sprintf("Date provided '%s' is not a valid RFC3339 date", r.GetDateAfter().GetValue()),
			)
		}

		query.DateAfter = &date
	}

	if r.GetLimit() != nil {
		v := r.GetLimit().GetValue()
		query.Limit = &v
	}

	if len(r.GetSeasonIds()) > 0 {
		query.SeasonIDs = r.GetSeasonIds()
	}

	if r.GetSort() != nil {
		v := r.GetSort().GetValue()
		query.SortBy = &v
	}

	if r.GetVenue() != nil {
		v := r.GetVenue().GetValue()
		query.Venue = &v
	}

	fixtures, err := s.fixtureRepo.ByTeamID(r.GetTeamId(), query)

	if err != nil {
		s.logger.Warnf("Error retrieving Fixture(s) in Result Service. Error: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	return s.sendResults(fixtures, stream)
}

func (s ResultService) GetResultsForSeason(r *statistico.SeasonRequest, stream statistico.ResultService_GetResultsForSeasonServer) error {
	date, err := time.Parse(time.RFC3339, r.DateBefore)

	if err != nil {
		return status.Error(codes.InvalidArgument, "Date provided is not a valid RFC3339 date")
	}

	id := uint64(r.SeasonId)

	query := app.FixtureRepositoryQuery{
		SeasonIDs: []uint64{id},
		DateTo:   &date,
	}

	fixtures, err := s.fixtureRepo.Get(query)

	if err != nil {
		s.logger.Warnf("Error retrieving Fixture(s) in Result Service. Error: %s", err.Error())
		return status.Error(codes.Internal, "Internal server error")
	}

	return s.sendResults(fixtures, stream)
}

func (s ResultService) sendResults(f []app.Fixture, stream statistico.ResultService_GetResultsForTeamServer) error {
	for _, fix := range f {
		x, err := s.factory.BuildResult(&fix)

		if err != nil {
			s.logger.Warnf("Error hydrating Result. Error: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		if err := stream.Send(x); err != nil {
			s.logger.Warnf("Error streaming Result back to client. Error: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}
	}

	return nil
}

func NewResultService(r app.FixtureRepository, f *factory.ResultFactory, log *logrus.Logger) *ResultService {
	return &ResultService{fixtureRepo: r, factory: f, logger: log}
}
