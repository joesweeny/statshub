package grpc

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/grpc/factory"
	"github.com/statistico/statistico-data/internal/app/grpc/proto"
	"time"
)

type FixtureService struct {
	fixtureRepo app.FixtureRepository
	factory     *factory.FixtureFactory
	logger      *logrus.Logger
}

func (s *FixtureService) ListSeasonFixtures(r *proto.SeasonFixtureRequest, stream proto.FixtureService_ListSeasonFixturesServer) error {
	from, err := time.Parse(time.RFC3339, r.DateFrom)

	if err != nil {
		return ErrTimeParse
	}

	to, err := time.Parse(time.RFC3339, r.DateTo)

	if err != nil {
		return ErrTimeParse
	}

	query := app.FixtureRepositoryQuery{
		DateTo:   &to,
		DateFrom: &from,
	}

	fixtures, err := s.fixtureRepo.Get(query)

	if err != nil {
		s.logger.Warnf("Error retrieving Fixture(s). Error: %s", err.Error())
		return internalServerError
	}

	for _, fix := range fixtures {
		f, err := s.factory.BuildFixture(&fix)

		if err != nil {
			s.logger.Warnf("Error hydrating Fixture. Error: %s", err.Error())
			continue
		}

		if err := stream.Send(f); err != nil {
			s.logger.Warnf("Error streaming Fixture back to client. Error: %s", err.Error())
			continue
		}
	}

	return nil
}

func (s *FixtureService) FixtureByID(c context.Context, r *proto.FixtureRequest) (*proto.Fixture, error) {
	fix, err := s.fixtureRepo.ByID(uint64(r.FixtureId))

	if err != nil {
		return nil, fmt.Errorf("fixture with ID %d does not exist", r.FixtureId)
	}

	f, err := s.factory.BuildFixture(fix)

	if err != nil {
		s.logger.Warnf("Error hydrating Fixture: %s", err.Error())
		return nil, err
	}

	return f, nil
}

func NewFixtureService(r app.FixtureRepository, f *factory.FixtureFactory, log *logrus.Logger) *FixtureService {
	return &FixtureService{fixtureRepo: r, factory: f, logger: log}
}
