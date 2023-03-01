package user

// imports
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KaiserPhemi/redoc-dms/pkg/models"
	"github.com/gin-gonic/gin"
)

// add a new user
func AddUser(c *gin.Context) {
	newUser := models.User{}
	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	added := newUser.AddUser()
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
		"user":    added,
	})

}

// fetch a single user
func FetchUser(c *gin.Context) {
	fmt.Println("the id", c.Param("id"))
	userId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, _ := models.FetchUserById(userId)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "User details fetched.",
		"user":    user,
	})
}

// fetch all users
func FetchAllUsers(c *gin.Context) {
	allUsers := models.FetchAllUsers()
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "All users fetched.",
		"users":   allUsers,
	})
}

// delete a user
func DeleteUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	deletedUser := models.DeleteUser(userId)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "User details deleted.",
		"user":    deletedUser,
	})
}
