package model

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	Get(ctx context.Context, uuid uuid.UUID) (*User, error)
	SignUp(ctx context.Context, u *User) error
}

type UserRepository interface {
	FindByID(ctx context.Context, uuid uuid.UUID) (*User, error)
}
