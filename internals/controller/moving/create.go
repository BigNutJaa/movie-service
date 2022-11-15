package moving

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/BigNutJaa/movie-service/internals/model/moving"
	apiV1 "github.com/BigNutJaa/movie-service/pkg/api/v1"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.MovingCreateRequest) (*apiV1.MovingCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.moving.Create",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		MovieName: request.GetMovieName(),
		Date:      request.GetDate(),
		Time:      request.GetTime(),
		CinemaNo:  request.GetCinemaNo(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.MovingCreateResponse{Id: int32(id)}, nil
}
