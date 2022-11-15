package warehouse

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/BigNutJaa/movie-service/internals/model/warehouse"
	apiV1 "github.com/BigNutJaa/movie-service/pkg/api/v1"
	"github.com/sirupsen/logrus"
)

func (c *Controller) Create(ctx context.Context, request *apiV1.WarehouseCreateRequest) (*apiV1.WarehouseCreateResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.category.Create",
	)
	defer span.Finish()
	logrus.Info("warehouse")
	logrus.Error("warehouse")
	span.LogKV("Input Handler", request)
	id, err := c.service.Create(ctx, &model.Request{
		Name:   request.GetName(),
		Detail: request.GetDetail(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.WarehouseCreateResponse{Id: int32(id)}, nil
}
