package users

import (
	"context"

	"github.com/BigNutJaa/movie-service/internals/entity"
	model "github.com/BigNutJaa/movie-service/internals/model/users"
)

func (s *UsersService) Get(ctx context.Context, request *model.FitterReadUsers) (*model.ReadResponseUsers, error) {
	makeFilter := s.makeFilterUsers(request)
	users := &entity.Users{}

	err := s.repository.Find(makeFilter, users)

	return &model.ReadResponseUsers{
		Id:          int32(users.ID),
		FullName:    users.FullName,
		Address:     users.Address,
		PhoneNumber: users.PhoneNumber,
		Gender:      users.Gender,
	}, err
}

func (s *UsersService) makeFilterUsers(filters *model.FitterReadUsers) (output map[string]interface{}) {
	output = make(map[string]interface{})

	if len(filters.FullName) > 0 {
		output["full_name"] = filters.FullName
	}
	if filters.Id > 0 {
		output["id"] = filters.Id
	}
	return output

}
