package user

// imports
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// add a new user
func AddUser(c *gin.Context){
	return 
}

// fetch a single user
func FetchUser(c *gin.Context){}

// fetch all users
func FetchAllUsers(c *gin.Context){
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "All users fetched.",
	})
}

// update a single user
func UpdateUser(c *gin.Context){}

// delete a user
func DeleteUser(c *gin.Context){}