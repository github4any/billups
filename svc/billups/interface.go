package billups

import (
	"context"
)

type Service interface {
	Choices(ctx context.Context) (interface{}, error)
	Choice(ctx context.Context) (interface{}, error)
	Play(ctx context.Context, req PlayerRequest) (interface{}, error)
	Scoreboard(ctx context.Context) (interface{}, error)
	Reset(ctx context.Context) error
}
