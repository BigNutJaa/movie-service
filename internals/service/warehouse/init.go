package warehouse

import (
	"github.com/BigNutJaa/movie-service/internals/repository/postgres"
)

type WarehouseService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &WarehouseService{
		repository: r,
	}
}
