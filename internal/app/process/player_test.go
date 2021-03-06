package process_test

import (
	"errors"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/statistico/statistico-data/internal/app"
	"github.com/statistico/statistico-data/internal/app/mock"
	"github.com/statistico/statistico-data/internal/app/process"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPlayerProcessor_Process(t *testing.T) {
	t.Run("inserts player if does not already exist", func(t *testing.T) {
		t.Helper()

		playerRepo := new(mock.PlayerRepository)
		squadRepo := new(mock.SquadRepository)
		requester := new(mock.PlayerRequester)
		//logger, hook := test.NewNullLogger()
		logger, _ := test.NewNullLogger()

		processor := process.NewPlayerProcessor(playerRepo, squadRepo, requester, logger)

		done := make(chan bool)

		def := newPlayer(1)
		mid := newPlayer(2)
		str := newPlayer(3)

		squad := newPlayerSquad()

		squadRepo.On("All").Return(squad, nil)

		playerRepo.On("ByID", uint64(1)).Return(&app.Player{}, errors.New("not found"))
		playerRepo.On("ByID", uint64(2)).Return(&app.Player{}, errors.New("not found"))
		playerRepo.On("ByID", uint64(3)).Return(&app.Player{}, errors.New("not found"))

		requester.On("PlayerByID", uint64(1)).Return(def, nil)
		requester.On("PlayerByID", uint64(2)).Return(mid, nil)
		requester.On("PlayerByID", uint64(3)).Return(str, nil)

		playerRepo.On("Insert", def).Return(nil)
		playerRepo.On("Insert", mid).Return(nil)
		playerRepo.On("Insert", str).Return(nil)

		processor.Process("player", "", done)

		<-done

		requester.AssertExpectations(t)
		playerRepo.AssertExpectations(t)
		squadRepo.AssertExpectations(t)
		//assert.Nil(t, hook.LastEntry())
	})

	t.Run("does not insert player if already exists", func(t *testing.T) {
		t.Helper()

		playerRepo := new(mock.PlayerRepository)
		squadRepo := new(mock.SquadRepository)
		requester := new(mock.PlayerRequester)
		//logger, hook := test.NewNullLogger()
		logger, _ := test.NewNullLogger()

		processor := process.NewPlayerProcessor(playerRepo, squadRepo, requester, logger)

		done := make(chan bool)

		def := newPlayer(1)
		mid := newPlayer(2)
		str := newPlayer(3)

		squad := newPlayerSquad()

		squadRepo.On("All").Return(squad, nil)

		playerRepo.On("ByID", uint64(1)).Return(&app.Player{}, errors.New("not found"))
		playerRepo.On("ByID", uint64(2)).Return(&app.Player{}, nil)
		playerRepo.On("ByID", uint64(3)).Return(&app.Player{}, errors.New("not found"))

		requester.On("PlayerByID", uint64(1)).Return(def, nil)
		requester.On("PlayerByID", uint64(3)).Return(str, nil)

		requester.AssertNotCalled(t, "PlayerByID", uint64(2))

		playerRepo.On("Insert", def).Return(nil)
		playerRepo.On("Insert", str).Return(nil)

		playerRepo.AssertNotCalled(t, "Insert", mid)

		processor.Process("player", "", done)

		<-done

		requester.AssertExpectations(t)
		playerRepo.AssertExpectations(t)
		squadRepo.AssertExpectations(t)
		//assert.Nil(t, hook.LastEntry())
	})

	t.Run("logs error if cannot insert player into repository", func(t *testing.T) {
		t.Helper()

		playerRepo := new(mock.PlayerRepository)
		squadRepo := new(mock.SquadRepository)
		requester := new(mock.PlayerRequester)
		logger, hook := test.NewNullLogger()

		processor := process.NewPlayerProcessor(playerRepo, squadRepo, requester, logger)

		done := make(chan bool)

		def := newPlayer(1)
		mid := newPlayer(2)
		str := newPlayer(3)

		squad := newPlayerSquad()

		squadRepo.On("All").Return(squad, nil)

		playerRepo.On("ByID", uint64(1)).Return(&app.Player{}, errors.New("not found"))
		playerRepo.On("ByID", uint64(2)).Return(&app.Player{}, nil)
		playerRepo.On("ByID", uint64(3)).Return(&app.Player{}, errors.New("not found"))

		requester.On("PlayerByID", uint64(1)).Return(def, nil)
		requester.On("PlayerByID", uint64(3)).Return(str, nil)

		requester.AssertNotCalled(t, "PlayerByID", uint64(2))

		playerRepo.On("Insert", def).Return(nil)
		playerRepo.On("Insert", str).Return(errors.New("cannot insert"))

		playerRepo.AssertNotCalled(t, "Insert", mid)

		processor.Process("player", "", done)

		<-done

		requester.AssertExpectations(t)
		playerRepo.AssertExpectations(t)
		squadRepo.AssertExpectations(t)
		assert.NotNil(t, hook.LastEntry().Message)
	})

	t.Run("player is not inserted if requester returns nil", func(t *testing.T) {
		t.Helper()

		playerRepo := new(mock.PlayerRepository)
		squadRepo := new(mock.SquadRepository)
		requester := new(mock.PlayerRequester)
		//logger, hook := test.NewNullLogger()
		logger, _ := test.NewNullLogger()

		processor := process.NewPlayerProcessor(playerRepo, squadRepo, requester, logger)

		done := make(chan bool)

		def := newPlayer(1)
		mid := newPlayer(2)
		str := newPlayer(3)

		squad := newPlayerSquad()

		squadRepo.On("All").Return(squad, nil)

		playerRepo.On("ByID", uint64(1)).Return(&app.Player{}, errors.New("not found"))
		playerRepo.On("ByID", uint64(2)).Return(&app.Player{}, errors.New("not found"))
		playerRepo.On("ByID", uint64(3)).Return(&app.Player{}, errors.New("not found"))

		requester.On("PlayerByID", uint64(1)).Return(def, nil)
		requester.On("PlayerByID", uint64(2)).Return(&app.Player{}, errors.New("oh damn"))
		requester.On("PlayerByID", uint64(3)).Return(str, nil)

		playerRepo.On("Insert", def).Return(nil)
		playerRepo.On("Insert", str).Return(nil)

		playerRepo.AssertNotCalled(t, "Insert", mid)

		processor.Process("player", "", done)

		<-done

		requester.AssertExpectations(t)
		playerRepo.AssertExpectations(t)
		squadRepo.AssertExpectations(t)
	})
}

func newPlayer(id uint64) *app.Player {
	return &app.Player{
		ID:        id,
		CountryId: uint64(154),
		FirstName: "Manuel",
		LastName:  "Lanzini",
	}
}

func newPlayerSquad() []app.Squad {
	var squads []app.Squad

	s := app.Squad{
		SeasonID:  45,
		TeamID:    98,
		PlayerIDs: []uint64{1, 2, 3},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	squads = append(squads, s)

	return squads
}
