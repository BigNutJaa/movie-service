package wrapper

import (
	"context"
	model "github.com/BigNutJaa/movie-service/internals/model/delta"

	"github.com/opentracing/opentracing-go"
)

func (wrp *Wrapper) Create2(ctx context.Context, input *model.Request) (int, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "Service.Delta2.Create2")
	defer sp.Finish()

	id, err := wrp.Service.Create(ctx, input)

	sp.LogKV("ID", id)
	sp.LogKV("err", err)

	return id, err
}
