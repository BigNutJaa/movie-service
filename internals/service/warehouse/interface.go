package warehouse

import (
	"context"
	model "github.com/BigNutJaa/movie-service/internals/model/warehouse"
)

//go:generate mockery --name=Service
type Service interface {
	Create(ctx context.Context, input *model.Request) (ID int, err error)
	Delete(ctx context.Context, request *model.FitterDeleteWarehouse) (*model.DeleteResponseWarehouse, error)
}
