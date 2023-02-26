package main

// imports
import (
	"net/http"

	"github.com/KaiserPhemi/redoc-dms/pkg/routes"
	"github.com/gin-gonic/gin"
)

// index route handler
func appMain(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Welcome to Redoc DMS.",
	})
}

// main function
func main() {
	router := gin.Default()
	router.GET("/", appMain)
	routes.RegisterUserRoutes()
	router.Run("localhost:7777")
}
