package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonathanbs9/memoriz3r/model"
	"github.com/jonathanbs9/memoriz3r/model/apperrors"
)

// Me
func (h *Handler) Me(c *gin.Context) {
	/*c.JSON(http.StatusOK, gin.H{
		"message": "This is ME",
	})*/
	user, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request for unknown reason: %v \n", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return
	}

	uuid := user.(*model.User).UUID

	u, err := h.UserService.Get(c, uuid)

	if err != nil {
		log.Printf("Uanble to find user: %v \n %v", uuid, err)
		e := apperrors.NewNotFound("user", uuid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
