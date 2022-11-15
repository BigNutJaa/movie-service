package wrapper

import (
	"go.uber.org/dig"

	service "github.com/BigNutJaa/movie-service/internals/service/users"
)

type Wrapper struct {
	dig.In  `name:"wrapperUsers"`
	Service service.Service
}
