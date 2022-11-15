package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/movie-service/internals/service/movie"
)

type Wrapper struct {
	dig.In  `name:"wrapperMovie"`
	Service service.Service
}
