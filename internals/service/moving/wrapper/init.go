package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/movie-service/internals/service/moving"
)

type Wrapper struct {
	dig.In  `name:"wrapperMoving"`
	Service service.Service
}
