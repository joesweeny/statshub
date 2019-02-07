package squad

import (
	"github.com/joesweeny/statshub/internal/season"
	"github.com/joesweeny/sportmonks-go-client"
	"log"
)

const callLimit = 1800
const squad = "squad"
const squadCurrentSeason = "squad:current-season"

type Service struct {
	Repository
	SeasonRepo season.Repository
	Factory
	Client *sportmonks.Client
	Logger *log.Logger
}

var counter int

func (s Service) Process(command string, done chan bool) {
	switch command {
	case squad:
		go s.allSeasons(done)
	case squadCurrentSeason:
		go s.currentSeason(done)
	default:
		s.Logger.Fatalf("Command %s is not supported", command)
		return
	}
}

func (s Service) allSeasons(done chan bool) {
	ids, err := s.SeasonRepo.Ids()

	if err != nil {
		s.Logger.Fatalf("Error when retrieving Season IDs: %s", err.Error())
		return
	}

	go s.handleSeasons(ids, done, &counter)
}

func (s Service) currentSeason(done chan bool) {
	ids, err := s.SeasonRepo.CurrentSeasonIds()

	if err != nil {
		s.Logger.Fatalf("Error when retrieving Season IDs: %s", err.Error())
		return
	}

	go s.handleSeasons(ids, done, &counter)
}

func (s Service) handleSeasons(ids []int, done chan bool, c *int) {
	for _, id := range ids {
		res, err := s.Client.TeamsBySeasonId(id, []string{}, 5)

		if err != nil {
			s.Logger.Printf("Error when calling client. Message: %s\n", err.Error())
			done <- true
		}

		for _, t := range res.Data {
			if *c >= callLimit {
				s.Logger.Printf("Api call limited reached %d calls\n", *c)
				done <- true
			}

			s.handleTeam(id, t, c, done)
		}
	}

	done <- true
}

func (s Service) handleTeam(seasonId int, t sportmonks.Team, c *int, done chan bool) {
	if _, err := s.BySeasonAndTeam(seasonId, t.ID); err != ErrNotFound {
		return
	}

	res, err := s.Client.SquadBySeasonAndTeam(seasonId, t.ID, []string{}, 5)

	*c++

	if err != nil {
		s.Logger.Printf("Error when calling client. Message: %s", err.Error())
		done <- true
	}

	s.persistSquad(seasonId, t.ID, &res.Data)

	return
}

func (s Service) persistSquad(seasonId, teamId int, m *[]sportmonks.SquadPlayer) {
	_, err := s.BySeasonAndTeam(seasonId, teamId)

	if err == ErrNotFound {
		created := s.createSquad(seasonId, teamId, m)

		if err := s.Insert(created); err != nil {
			s.Logger.Printf("Error '%s' occurred when inserting Squad struct: %+v\n,", err.Error(), created)
		}

		return
	}

	return
}
