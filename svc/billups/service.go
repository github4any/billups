package billups

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billups/api/internal/rpssl"
	"github.com/billups/api/svc/billups/repository"
)

var (
	sets = [5]string{
		"rock",
		"paper",
		"scissors",
		"lizard",
		"spock",
	}
)

type (
	// Service struct
	service struct {
		billupsURL string
		repo       billupsRepository
	}

	billupsRepository interface {
		CreateScoreboard(ctx context.Context, arg repository.CreateScoreboardParams) error
		GetScoreboard(ctx context.Context) ([]repository.Scoreboard, error)
		DeleteScoreboard(ctx context.Context) error
	}
)

// NewService is a factory function,
// returns a new instance of the Service interface implementation
func NewService(repo billupsRepository, billupsURL string) Service {
	return &service{
		billupsURL: billupsURL,
		repo:       repo,
	}
}

func (s *service) Choices(ctx context.Context) (interface{}, error) {
	return []ChoiceResponse{
		{
			Id:   1,
			Name: sets[0],
		},
		{
			Id:   2,
			Name: sets[1],
		},
		{
			Id:   3,
			Name: sets[2],
		},
		{
			Id:   4,
			Name: sets[3],
		},
		{
			Id:   5,
			Name: sets[4],
		},
	}, nil
}

func (s *service) Choice(ctx context.Context) (interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", s.billupsURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var random RandomResponse
	if response.StatusCode == http.StatusOK {
		if err := json.NewDecoder(response.Body).Decode(&random); err != nil {
			return nil, err
		}
	}

	i := random.RandomNumber % 5
	switch i {
	case 0:
		return ChoiceResponse{
			Id:   5,
			Name: sets[4],
		}, nil
	case 1:
		return ChoiceResponse{
			Id:   1,
			Name: sets[0],
		}, nil
	case 2:
		return ChoiceResponse{
			Id:   2,
			Name: sets[1],
		}, nil
	case 3:
		return ChoiceResponse{
			Id:   3,
			Name: sets[2],
		}, nil
	case 4:
		return ChoiceResponse{
			Id:   4,
			Name: sets[3],
		}, nil
	}

	return nil, nil
}

func (s *service) Play(ctx context.Context, req PlayerRequest) (interface{}, error) {
	compAnswer := rpssl.GetRandomNumber()

	if err := s.repo.CreateScoreboard(ctx, repository.CreateScoreboardParams{
		Results:  rpssl.GetResult(sets[req.Player-1], sets[compAnswer]),
		Player:   int32(req.Player),
		Computer: int32(compAnswer + 1),
	}); err != nil {
		return nil, fmt.Errorf("save results to db, err: %w", err)
	}

	return PlayResponse{
		Results:  rpssl.GetResult(sets[req.Player-1], sets[compAnswer]),
		Player:   req.Player,
		Computer: compAnswer + 1,
	}, nil
}

func (s *service) Scoreboard(ctx context.Context) (interface{}, error) {
	result := make([]ScoreboardResponse, 0)

	scoreboard, err := s.repo.GetScoreboard(ctx)
	if err != nil {
		return nil, fmt.Errorf("scoreboard list, %w", err)
	}

	for _, s := range scoreboard {
		result = append(result, ScoreboardResponse{
			Results:   s.Results,
			Player:    int(s.Player),
			Computer:  int(s.Computer),
			Timestamp: s.CreatedAt,
		})
	}

	return result, nil
}

func (s *service) Reset(ctx context.Context) error {
	err := s.repo.DeleteScoreboard(ctx)
	if err != nil {
		return fmt.Errorf("scoreboard reset, %w", err)
	}

	return nil
}
