package mocks

import (
	"context"

	"github.com/jonathanbs9/memoriz3r/model"
	"github.com/stretchr/testify/mock"
)

type MockTokenService struct {
	mock.Mock
}

func (m *MockTokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenId string) (*model.TokenPair, error) {
	ret := m.Called(ctx, u, prevTokenId)

	// First value passed to "return"
	var r0 *model.TokenPair
	if ret.Get(0) != nil {
		// We can just return this if we know we won't  be passing function to "return"
		r0 = ret.Get(0).(*model.TokenPair)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}
	return r0, r1
}
