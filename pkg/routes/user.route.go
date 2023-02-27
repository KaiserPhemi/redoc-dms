package routes

// imports
import (
	user "github.com/KaiserPhemi/redoc-dms/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// users routes
func userRoutes(superRoute *gin.RouterGroup){
	usersRouter := superRoute.Group("/users")
	{
		usersRouter.GET("/", user.FetchAllUsers)
		usersRouter.POST("/", user.AddUser)
		usersRouter.DELETE("/:id", user.DeleteUser)
		usersRouter.PUT("/:id", user.FetchUser)
	}
}

