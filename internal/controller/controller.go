package controller

import (
	"github.com/radityarestan/mini-project/internal/service"
)

type Controller struct {
	service *service.Service
}

func NewController(service *service.Service) *Controller {
	return &Controller{service: service}
}
