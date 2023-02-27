package main

// imports
import (
	"github.com/KaiserPhemi/redoc-dms/pkg/routes"
	"github.com/gin-gonic/gin"
)

// main function
func main() {
	app := gin.New()
	router := app.Group("/api/v1")
	routes.AddRoutes(router)
	app.Run(":7777")
}
