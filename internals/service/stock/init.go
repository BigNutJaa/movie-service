package stock

import (
	"github.com/BigNutJaa/movie-service/internals/repository/postgres"
)

type StockService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &StockService{
		repository: r,
	}
}
