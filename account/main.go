package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting Server =) ")

	router := gin.Default()
	router.GET("/api/account", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Everything is OK",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	srv.ListenAndServe()
}
