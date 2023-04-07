package routes

// imports
import (
	"github.com/gofiber/fiber/v2"
)

// add all routes
func AddRoutes(superRoute *fiber.Group) {
	userRoutes(superRoute)
}
