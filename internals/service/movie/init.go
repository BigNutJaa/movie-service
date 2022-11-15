package movie

import (
	"github.com/BigNutJaa/movie-service/internals/repository/postgres"
)

type MovieService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &MovieService{
		repository: r,
	}
}
