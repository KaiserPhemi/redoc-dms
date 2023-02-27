package routes

// imports
import (
	"github.com/gin-gonic/gin"
)

// add all routes
func AddRoutes(superRoute *gin.RouterGroup) {
	userRoutes(superRoute)
}
