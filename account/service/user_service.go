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

// NewUserservice if a factory function for initializing a UserService
// with its repository layer dependencies
func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *UserService) Get(ctx context.Context, uuid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uuid)
	return u, err
}

// Sign Up reaches our to a UserRepository to verify the email address is available and signs up
// the user if this is the case
func (s *UserService) SignUp(ctx context.Context, u *model.User) error {
	panic("Method not implemented")
}
