package service

import "github.com/radityarestan/mini-project/internal/repository"

type Service struct {
	usersRepo *repository.UserRepository
}

func NewService(usersRepo *repository.UserRepository) *Service {
	return &Service{usersRepo: usersRepo}
}
