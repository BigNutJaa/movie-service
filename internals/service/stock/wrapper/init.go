package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/movie-service/internals/service/stock"
)

type Wrapper struct {
	dig.In  `name:"wrapperStock"`
	Service service.Service
}
