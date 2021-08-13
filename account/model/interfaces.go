package model

import "github.com/google/uuid"

type UserService interface {
	Get(uuid uuid.UUID) (*User, error)
}
