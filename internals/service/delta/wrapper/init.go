package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/movie-service/internals/service/delta"
)

type Wrapper struct {
	dig.In  `name:"wrapperDelta"`
	Service service.Service
}
