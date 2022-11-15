package category

import (
	"context"

	model "github.com/BigNutJaa/movie-service/internals/model/category"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
	Patch(ctx context.Context, request *model.FitterUpdateCategory) (*model.UpdateResponseCategory, error)
}
