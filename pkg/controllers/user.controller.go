package user

// imports
import (
	"strconv"

	"github.com/KaiserPhemi/redoc-dms/pkg/models"
	util "github.com/KaiserPhemi/redoc-dms/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func AddUser(c *fiber.Ctx) error {
	newUser := models.User{}
	if err := c.JSON(&newUser); err != nil {
		return c.Status(fiber.ErrBadGateway.Code).Send(nil)
	}
	createdUser := newUser.AddUser()
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Message": "User created successfully.",
		"User":    createdUser,
	})
}

// fetch a single user
func FetchUser(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	util.HandleError(err)
	user, _ := models.FetchUserById(userId)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User details fetched.",
		"user":    user,
	})
}

// fetch all users
func FetchAllUsers(c *fiber.Ctx) error {
	allUsers := models.FetchAllUsers()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "All users fetched.",
		"users":   allUsers,
	})
}

// delete a user
func DeleteUser(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	util.HandleError(err)
	deletedUser := models.DeleteUser(userId)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User details deleted.",
		"user":    deletedUser,
	})
}
