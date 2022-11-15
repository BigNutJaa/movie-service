package delta

import (
	"github.com/BigNutJaa/movie-service/internals/repository/postgres"
)

type DeltaService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &DeltaService{
		repository: r,
	}
}
