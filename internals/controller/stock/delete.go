package stock

import (
	"context"
	"github.com/opentracing/opentracing-go"
	model "github.com/BigNutJaa/movie-service/internals/model/stock"
	apiV1 "github.com/BigNutJaa/movie-service/pkg/api/v1"
)

func (c *Controller) Delete(ctx context.Context, request *apiV1.StockDeleteRequest) (*apiV1.StockDeleteResponse, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(
		ctx,
		opentracing.GlobalTracer(),
		"handler.stock.Delete",
	)
	defer span.Finish()
	span.LogKV("Input Handler", request)
	stockData, err := c.service.Delete(ctx, &model.FitterDeleteStock{
		Id: request.GetId(),
	})

	if err != nil {
		return nil, err
	}
	return &apiV1.StockDeleteResponse{
		Name:   stockData.Name,
		Detail: stockData.Detail,
		Qty:    stockData.Qty,
		Id:     stockData.Id,
	}, nil
}
