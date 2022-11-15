package users

import (
	"github.com/BigNutJaa/movie-service/internals/repository/postgres"
)

type UsersService struct {
	repository postgres.Repository
}

func NewService(r postgres.Repository) (service Service) {
	return &UsersService{
		repository: r,
	}
}
