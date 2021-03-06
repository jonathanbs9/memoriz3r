package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/jonathanbs9/memoriz3r/model"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindByID(ctx context.Context, uuid uuid.UUID) (*model.User, error) {
	ret := m.Called(ctx, uuid)

	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
