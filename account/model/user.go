package model

import (
	"github.com/google/uuid"
)

// User defines domain model and its json db representation
type User struct {
	UUID     uuid.UUID `db:"uuid" json:"uuid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"-"`
	Name     string    `db:"name" json:"name"`
	ImageURL string    `db:"img_url" json:"imgUrl"`
	Website  string    `db:"website" json:"website"`
}
