package routes

// imports
import (
	user "github.com/KaiserPhemi/redoc-dms/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// register all user routes
var RegisterUserRoutes = func() {
	router := gin.Default()
	router.GET("/users", user.FetchAllUsers)
	router.POST("/users", user.AddUser)
	router.GET("/users/{userId}", user.FetchUser)
	router.PUT("/users/{userId}", user.UpdateUser)
	router.DELETE("/users/{userId}", user.DeleteUser)
}
