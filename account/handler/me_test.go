package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jonathanbs9/memoriz3r/model"
	"github.com/jonathanbs9/memoriz3r/model/apperrors"
	"github.com/jonathanbs9/memoriz3r/model/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMe(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// SUCCESS
	t.Run("Success", func(t *testing.T) {

		uuid, _ := uuid.NewRandom()

		mockUserResp := &model.User{
			UUID:  uuid,
			Email: "jonathans@gmail.com",
			Name:  "Jonathan Brull Schreoder",
		}

		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uuid).Return(mockUserResp, nil)

		// A response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{
				UUID: uuid,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/api/account/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"user": mockUserResp,
		})
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t) // Assert that UserService.Get wall called
	})

	// NO CONTEXT USER
	t.Run("NoContextUser", func(t *testing.T) {
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, mock.Anything).Return(nil, nil)

		// Response Recorder for getting written http response
		rr := httptest.NewRecorder()

		// Do not append user to context
		router := gin.Default()
		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/api/account/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockUserService.AssertNotCalled(t, "Get", mock.Anything)
	})

	// NOT FOUNDGO
	t.Run("NotFound", func(t *testing.T) {
		uuid, _ := uuid.NewRandom()
		mockUserService := new(mocks.MockUserService)
		mockUserService.On("Get", mock.Anything, uuid).Return(nil, fmt.Sprintf("Some Error down call chain"))

		// Response Recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()
		router.Use(func(c *gin.Context) {
			c.Set("user", &model.User{
				UUID: uuid,
			})
		})

		NewHandler(&Config{
			R:           router,
			UserService: mockUserService,
		})

		request, err := http.NewRequest(http.MethodGet, "/api/account/me", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)
		respErr := apperrors.NewNotFound("user", uuid.String())

		respBody, err := json.Marshal(gin.H{
			"error": respErr,
		})
		assert.NoError(t, err)
		log.Println("RR CODE")
		log.Println(rr.Code)
		log.Println(respErr.Status())
		log.Println("RESPERROR CODE")
		assert.Equal(t, respErr.Status(), rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockUserService.AssertExpectations(t)

	})
}
