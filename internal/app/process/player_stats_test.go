package process_test

import (
	"errors"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/mock"
	"github.com/statistico/statistico-data/internal/app/process"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPlayerStatsProcessor_Process(t *testing.T) {
	t.Run("inserts new player stats into repository when processing player stats by season id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		requester.On("PlayerStatsBySeasonIDs", []uint64{45}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("Insert", one).Return(nil)
		playerStatsRepo.On("Insert", two).Return(nil)

		processor.Process("player-stats:by-season-id", "45", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error if unable to insert player stats into repository when processing player stats by season id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		requester.On("PlayerStatsBySeasonIDs", []uint64{45}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("Insert", one).Return(errors.New("error occurred"))
		playerStatsRepo.On("Insert", two).Return(nil)

		processor.Process("player-stats:by-season-id", "45", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("update existing player stats into repository when processing player stats by season id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		requester.On("PlayerStatsBySeasonIDs", []uint64{45}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(one, nil)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(two, nil)
		playerStatsRepo.On("Update", one).Return(nil)
		playerStatsRepo.On("Update", two).Return(nil)

		processor.Process("player-stats:by-season-id", "45", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error if unable to update player stats into repository when processing player stats by season id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		requester.On("PlayerStatsBySeasonIDs", []uint64{45}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(one, nil)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(two, nil)
		playerStatsRepo.On("Update", one).Return(nil)
		playerStatsRepo.On("Update", two).Return(errors.New("error occurred"))

		processor.Process("player-stats:by-season-id", "45", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("inserts new player stats into repository when processing player stats by date command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		competitionRepo.On("IDs").Return([]uint64{1, 2, 3}, nil)

		date, _ := time.Parse("2006-01-02", "2021-01-18")

		requester.On("PlayerStatsByDate", date, []uint64{1, 2, 3}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("Insert", one).Return(nil)
		playerStatsRepo.On("Insert", two).Return(nil)

		processor.Process("player-stats:by-date", "2021-01-18", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error if unable to insert player stats into repository when processing player stats by date command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		competitionRepo.On("IDs").Return([]uint64{1, 2, 3}, nil)

		date, _ := time.Parse("2006-01-02", "2021-01-18")

		requester.On("PlayerStatsByDate", date, []uint64{1, 2, 3}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("Insert", one).Return(errors.New("error occurred"))
		playerStatsRepo.On("Insert", two).Return(nil)

		processor.Process("player-stats:by-date", "2021-01-18", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("update existing player stats into repository when processing player stats by date command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		competitionRepo.On("IDs").Return([]uint64{1, 2, 3}, nil)

		date, _ := time.Parse("2006-01-02", "2021-01-18")

		requester.On("PlayerStatsByDate", date, []uint64{1, 2, 3}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(one, nil)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(two, nil)
		playerStatsRepo.On("Update", one).Return(nil)
		playerStatsRepo.On("Update", two).Return(nil)

		processor.Process("player-stats:by-date", "2021-01-18", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error if unable to update player stats into repository when processing player stats by date command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		competitionRepo.On("IDs").Return([]uint64{1, 2, 3}, nil)

		date, _ := time.Parse("2006-01-02", "2021-01-18")

		requester.On("PlayerStatsByDate", date, []uint64{1, 2, 3}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(one, nil)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(two, nil)
		playerStatsRepo.On("Update", one).Return(nil)
		playerStatsRepo.On("Update", two).Return(errors.New("error occurred"))

		processor.Process("player-stats:by-date", "2021-01-18", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		seasonRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("inserts new player stats into repository when processing player stats by competition id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("PlayerStatsBySeasonIDs", []uint64{1, 2}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("Insert", one).Return(nil)
		playerStatsRepo.On("Insert", two).Return(nil)

		processor.Process("player-stats:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error if unable to insert player stats into repository when processing player stats by competition id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("PlayerStatsBySeasonIDs", []uint64{1, 2}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(&app.PlayerStats{}, errors.New("not found"))
		playerStatsRepo.On("Insert", one).Return(errors.New("error occurred"))
		playerStatsRepo.On("Insert", two).Return(nil)

		processor.Process("player-stats:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})

	t.Run("update existing player stats into repository when processing player stats by competition id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("PlayerStatsBySeasonIDs", []uint64{1, 2}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(one, nil)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(two, nil)
		playerStatsRepo.On("Update", one).Return(nil)
		playerStatsRepo.On("Update", two).Return(nil)

		processor.Process("player-stats:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error if unable to update player stats into repository when processing player stats by competition id command", func(t *testing.T) {
		t.Helper()

		playerStatsRepo := new(mock.PlayerStatsRepository)
		competitionRepo := new(mock.CompetitionRepository)
		seasonRepo := new(mock.SeasonRepository)
		requester := new(mock.PlayerStatsRequester)
		clock := clockwork.NewFakeClock()
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerStatsProcessor(playerStatsRepo, competitionRepo, seasonRepo, requester, clock, logger)

		done := make(chan bool)

		one := newPlayerStats(45, 99)
		two := newPlayerStats(45, 5)

		stats := make([]*app.PlayerStats, 2)
		stats[0] = one
		stats[1] = two

		ch := playerStatsChannel(stats)

		seasonRepo.On("ByCompetitionId", uint64(5), "name_asc").Return(
			[]app.Season{
				*newSeason(1, false),
				*newSeason(2, true),
			},
			nil,
		)

		requester.On("PlayerStatsBySeasonIDs", []uint64{1, 2}).Return(ch)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(99)).Return(&app.PlayerStats{}, nil)
		playerStatsRepo.On("ByFixtureAndPlayer", uint64(45), uint64(5)).Return(&app.PlayerStats{}, nil)
		playerStatsRepo.On("Update", one).Return(nil)
		playerStatsRepo.On("Update", two).Return(errors.New("error occurred"))

		processor.Process("player-stats:by-competition-id", "5", done)

		<-done

		requester.AssertExpectations(t)
		playerStatsRepo.AssertExpectations(t)
		competitionRepo.AssertExpectations(t)
		assert.Equal(t, 1, len(hook.Entries))
		assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	})
}

func newPlayerStats(fixtureID, playerID uint64) *app.PlayerStats {
	pos := "M"
	return &app.PlayerStats{
		FixtureID:       fixtureID,
		PlayerID:        playerID,
		TeamID:          uint64(44),
		Position:        &pos,
		IsSubstitute:    false,
		PlayerShots:     app.PlayerShots{},
		PlayerGoals:     app.PlayerGoals{},
		PlayerFouls:     app.PlayerFouls{},
		PlayerCrosses:   app.PlayerCrosses{},
		PlayerPasses:    app.PlayerPasses{},
		PlayerPenalties: app.PlayerPenalties{},
	}
}

func playerStatsChannel(stats []*app.PlayerStats) chan *app.PlayerStats {
	ch := make(chan *app.PlayerStats, len(stats))

	for _, c := range stats {
		ch <- c
	}

	close(ch)

	return ch
}
