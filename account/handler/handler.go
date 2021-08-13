package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

type Config struct {
	R *gin.Engine
	//US *model.UserService
}

func NewHandler(c *Config) {
	h := &Handler{}
	//os.Getenv("ACCOUNT_API_URL"))
	g := c.R.Group("api/account/")

	g.GET("/me", h.Me)
	g.POST("/signup", h.SignUp)
	g.POST("/signin", h.SignIn)
	g.POST("/signout", h.SignOut)
	g.POST("/tokens", h.Token)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

	/*g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Everything will be ok",
		})
	})*/
}

// Me
func (h *Handler) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is ME",
	})
}

// Signup
func (h *Handler) SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SignUp",
	})
}

// Signin
func (h *Handler) SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SignIn",
	})
}

// Signout
func (h *Handler) SignOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Bye SignOut",
	})
}

// Tokens
func (h *Handler) Token(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Tokenssss ",
	})
}

// Signout
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Image Handler",
	})
}

// Delete
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Image",
	})
}

// Delete
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "This is details",
	})
}
