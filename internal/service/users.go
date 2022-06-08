package service

import (
	"database/sql"

	"github.com/radityarestan/mini-project/internal/shared"
	"github.com/radityarestan/mini-project/internal/shared/dto"
)

type User interface {
	Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

func (s *Service) Register(req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	oldUser, err := s.usersRepo.FetchUserByUsername(req.Username)

	if err != nil && err != sql.ErrNoRows {
		return &dto.RegisterResponse{}, err
	}

	if oldUser != nil {
		return &dto.RegisterResponse{}, shared.ErrUserAlreadyExist
	}

	user, err := s.usersRepo.InsertUser(
		req.Name, req.Username, req.Password, req.Role,
	)

	if err != nil {
		return &dto.RegisterResponse{}, err
	}

	return &dto.RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
