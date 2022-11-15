package alpha

import (
	"context"
	"github.com/BigNutJaa/movie-service/internals/entity"
	model "github.com/BigNutJaa/movie-service/internals/model/alpha"
)

func (s *AlphaService) Create2(ctx context.Context, request *model.Request) (int, error) {
	//defer wg.Done()

	input := &entity.Alpha{
		MovieName: request.MovieName,
		Date:      request.Date,
		Time:      request.Time,
		CinemaNo:  request.CinemaNo,
	}

	err := s.repository.Create(input)

	return input.ID, err
}
