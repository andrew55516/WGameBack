package main

import (
	"WGameBack/initializers"
	routes "WGameBack/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	initializers.LoadEnvVar()
}

func main() {
	r := gin.Default()

	routes.AuthRoutes(r)
	routes.UserRoutes(r)

	r.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "access granted for api-1"})
	})

	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "access granted for api-2"})
	})

	r.Run()
}
