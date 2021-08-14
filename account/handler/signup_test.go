package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jonathanbs9/memoriz3r/model"
	"github.com/jonathanbs9/memoriz3r/model/apperrors"
	"github.com/jonathanbs9/memoriz3r/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUp(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Email and password required", func(t *testing.T) {
		mockUserservice := new(mocks.MockUserService)
		// If we call de Signup method => return nil
		mockUserservice.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// A response recorder for getting written http response
		rr := httptest.NewRecorder()

		// Don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserservice,
		})

		// Create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email": "",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create  a reader
		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUserservice.AssertNotCalled(t, "SignUp")
	})

	t.Run("Invalid Email", func(t *testing.T) {
		mockUserservice := new(mocks.MockUserService)
		// If we call de Signup method => return nil
		mockUserservice.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// A response recorder for getting written http response
		rr := httptest.NewRecorder()

		// Don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserservice,
		})

		// Create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "sin@com",
			"password": "validPassword",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create  a reader
		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUserservice.AssertNotCalled(t, "SignUp")
	})

	t.Run("Password too short", func(t *testing.T) {
		mockUserservice := new(mocks.MockUserService)
		// If we call de Signup method => return nil
		mockUserservice.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// A response recorder for getting written http response
		rr := httptest.NewRecorder()

		// Don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserservice,
		})

		// Create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "valid@email.com",
			"password": "abc",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create  a reader
		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUserservice.AssertNotCalled(t, "SignUp")
	})

	t.Run("Password too long", func(t *testing.T) {
		mockUserservice := new(mocks.MockUserService)
		// If we call de Signup method => return nil
		mockUserservice.On("SignUp", mock.AnythingOfType("*gin.Context"), mock.AnythingOfType("*model.User")).Return(nil)

		// A response recorder for getting written http response
		rr := httptest.NewRecorder()

		// Don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserservice,
		})

		// Create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    "valid@email.com",
			"password": "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create  a reader
		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUserservice.AssertNotCalled(t, "SignUp")
	})
	// ** Chequear porque devuelve code 200
	t.Run("Error returned from UserService", func(t *testing.T) {
		u := &model.User{
			Email:    "valid@email.com",
			Password: "validpassword",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(apperrors.NewConflict("User already exists ", u.Email))

		// A response recorder for getting written http response
		rr := httptest.NewRecorder()

		// Don't need a middleware as we don't yet have  authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		// Create a request body with ..
		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		assert.Equal(t, 409, rr.Code)
		mockUserService.AssertExpectations(t)
	})

	t.Run("Successful TOKEN Created", func(t *testing.T) {
		u := &model.User{
			Email:    "valid@email.com",
			Password: "AvalidPassword123",
		}

		mockTokenResp := &model.TokenPair{
			IDToken:      "IDToken",
			RefreshToken: "refreshToken",
		}

		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").Return(mockTokenResp, nil)

		// Response Recorder for getting written http response
		rr := httptest.NewRecorder()

		// Don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		// Create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		// User bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"tokens": mockTokenResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})

	t.Run("Failed token Creation", func(t *testing.T) {
		u := &model.User{
			Email:    "jonathan@email.com",
			Password: "thisPassword123",
		}

		mockErrorResponse := apperrors.NewInternal()
		mockUserService := new(mocks.MockUserService)
		mockTokenService := new(mocks.MockTokenService)

		mockUserService.On("SignUp", mock.AnythingOfType("*gin.Context"), u).Return(nil)
		mockTokenService.On("NewPairFromUser", mock.AnythingOfType("*gin.Context"), u, "").Return(nil, mockErrorResponse)

		// A response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:            router,
			UserService:  mockUserService,
			TokenService: mockTokenService,
		})

		reqBody, err := json.Marshal(gin.H{
			"email":    u.Email,
			"password": u.Password,
		})
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/api/account/signup", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"error": mockErrorResponse,
		})
		assert.NoError(t, err)

		assert.Equal(t, mockErrorResponse.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())

		mockUserService.AssertExpectations(t)
		mockTokenService.AssertExpectations(t)
	})
}
