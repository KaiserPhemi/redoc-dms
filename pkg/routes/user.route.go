package routes

// imports
import (
	user "github.com/KaiserPhemi/redoc-dms/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

// users routes
func userRoutes(superRoute *fiber.Group){
	usersRouter := superRoute.Group("/users")
	{
		usersRouter.Get("/", user.FetchAllUsers)
		usersRouter.Post("/", user.AddUser)
		usersRouter.Delete("/:id", user.DeleteUser)
		usersRouter.Get("/:id", user.FetchUser)
	}
}

