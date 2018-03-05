package api

import (
	"net/http"

	"./price"
	"github.com/gin-gonic/gin"
)

// Run starts up the API for commonwand
func Run() {
	router := gin.Default()
	message := "sup"

	router.POST("rest/price", price.HandlePost)

	router.GET("rest/prices", func(c *gin.Context) {
		c.String(http.StatusOK, message)
	})
	router.GET("rest/price/:id", func(c *gin.Context) {
		c.String(http.StatusOK, message)
	})
	router.PUT("rest/price/:id", func(c *gin.Context) {
		c.String(http.StatusOK, message)
	})
	router.POST("users/login", func(c *gin.Context) {
		c.String(http.StatusOK, message)
	})
	router.POST("users/signup", func(c *gin.Context) {
		c.String(http.StatusOK, message)
	})
	router.POST("users/me	", func(c *gin.Context) {
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
