package main

// imports
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// index route handler
func appMain(c *gin.Context) {

	resp := "Welcome to redoc document management system."
	c.IndentedJSON(http.StatusOK, resp)

}

func main() {
	router := gin.Default()
	router.GET("/", appMain)
	router.Run("localhost:7777")
}
