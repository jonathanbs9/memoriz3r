package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jonathanbs9/memoriz3r/model"
	"github.com/jonathanbs9/memoriz3r/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		uuid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UUID:  uuid,
			Email: "jonathan@test.com.ar",
			Name:  "Jonathan",
		}

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})

		mockUserRepository.On("FindByID", mock.Anything, uuid).Return(mockUserResp, nil)

		ctx := context.TODO()
		u, err := us.Get(ctx, uuid)

		assert.NoError(t, err)
		assert.Equal(t, u, mockUserResp)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		uuid, _ := uuid.NewRandom()

		mockUserRepository := new(mocks.MockUserRepository)
		us := NewUserService(&USConfig{
			UserRepository: mockUserRepository,
		})

		mockUserRepository.On("FindByID", mock.Anything, uuid).Return(nil, fmt.Errorf("Some error!"))

		ctx := context.TODO()
		u, err := us.Get(ctx, uuid)

		assert.Nil(t, u)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}
