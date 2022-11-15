package users

import (
	"github.com/BigNutJaa/movie-service/internals/service/users/wrapper"
)

type Controller struct {
	service wrapper.Wrapper
}

func NewController(s wrapper.Wrapper) *Controller {
	return &Controller{
		service: s,
	}
}
