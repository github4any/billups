package billups

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
)

type (
	// Endpoints collection of billup service
	Endpoints struct {
		Choices    endpoint.Endpoint
		Choice     endpoint.Endpoint
		Play       endpoint.Endpoint
		Scoreboard endpoint.Endpoint
		Reset      endpoint.Endpoint
	}

	PlayerRequest struct {
		Player int `json:"player"`
	}

	ChoiceResponse struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	RandomResponse struct {
		RandomNumber int `json:"random_number"`
	}

	PlayResponse struct {
		Results  string `json:"results"`
		Player   int    `json:"player"`
		Computer int    `json:"computer"`
	}

	ScoreboardResponse struct {
		Results   string    `json:"results"`
		Player    int       `json:"player"`
		Computer  int       `json:"computer"`
		Timestamp time.Time `json:"timestamp"`
	}
)

func MakeEndpoints(s Service, m ...endpoint.Middleware) Endpoints {
	e := Endpoints{
		Choices:    MakeChoicesEndpoint(s),
		Choice:     MakeChoiceEndpoint(s),
		Play:       MakePlayEndpoint(s),
		Scoreboard: MakeScoreboardEndpoint(s),
		Reset:      MakeResetEndpoint(s),
	}

	// setup middlewares for each endpoints
	if len(m) > 0 {
		for _, mdw := range m {
			e.Choices = mdw(e.Choices)
			e.Choice = mdw(e.Choice)
			e.Play = mdw(e.Play)
			e.Scoreboard = mdw(e.Scoreboard)
			e.Reset = mdw(e.Reset)
		}
	}

	return e
}

// MakeEndpoint ...
func MakeChoicesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return s.Choices(ctx)
	}
}

func MakeChoiceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return s.Choice(ctx)
	}
}

func MakePlayEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(PlayerRequest)

		return s.Play(ctx, req)
	}
}

func MakeScoreboardEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		return s.Scoreboard(ctx)
	}
}

func MakeResetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		if err := s.Reset(ctx); err != nil {
			return nil, err
		}

		return true, nil
	}
}
