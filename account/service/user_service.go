package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jonathanbs9/memoriz3r/model"
)

type UserService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

func (s *UserService) Get(ctx context.Context, uuid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uuid)
	return u, err
}
