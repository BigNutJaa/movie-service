package wrapper

import (
	"github.com/BigNutJaa/movie-service/internals/repository/postgres"
)

type ProductService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &ProductService{
		repository: r,
	}
}
