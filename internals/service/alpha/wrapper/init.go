package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/movie-service/internals/service/alpha"
)

type Wrapper struct {
	dig.In  `name:"wrapperAlpha"`
	Service service.Service
}
