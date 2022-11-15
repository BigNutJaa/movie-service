package delta

import (
	"context"
	model "github.com/BigNutJaa/movie-service/internals/model/delta"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
}
