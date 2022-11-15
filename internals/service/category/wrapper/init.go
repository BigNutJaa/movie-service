package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/movie-service/internals/service/category"
)

type Wrapper struct {
	dig.In  `name:"wrapperCategory"`
	Service service.Service
}
