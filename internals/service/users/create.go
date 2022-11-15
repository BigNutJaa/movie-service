package users

import (
	"context"

	"github.com/BigNutJaa/movie-service/internals/entity"
	model "github.com/BigNutJaa/movie-service/internals/model/users"
)

func (s *UsersService) Create(ctx context.Context, request *model.Request) (int, error) {

	input := &entity.Users{
		FullName:    request.FirstName + " " + request.LastName,
		Address:     request.Address,
		PhoneNumber: request.PhoneNumber,
		Gender:      request.Gender,
	}

	err := s.repository.Create(input)

	//sp.LogKV("Repository result  :", err)

	return input.ID, err
}
