package controller

import (
	"github.com/radityarestan/mini-project/repository"
)

type Controller struct {
	usersRepo *repository.UserRepository
}

func NewController(usersRepo *repository.UserRepository) *Controller {
	return &Controller{usersRepo: usersRepo}
}
