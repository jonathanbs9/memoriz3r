package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathanbs9/memoriz3r/model"
	"github.com/jonathanbs9/memoriz3r/model/apperrors"
)

type signupreq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup
func (h *Handler) SignUp(c *gin.Context) {
	/*c.JSON(http.StatusOK, gin.H{
		"message": "SignUp",
	})*/
	var req signupreq

	// Bind incoming json to struct and check for validation errors
	if ok := bindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.UserService.SignUp(c, u)

	if err != nil {
		log.Printf("Failed to sign up user %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// Create token pair as strings
	tokens, err := h.TokenService.NewPairFromUser(c, u, "")
	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		// May eventually implement rollback logic here
		// meaning, if we fail to create tokens after creating a user,
		// we make sure to clear/delete the created user in the database

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
