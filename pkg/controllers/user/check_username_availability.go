package userController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func CheckUsernameAvailability(c *gin.Context) {
	username := c.Params.ByName("username")

	user, err := userModel.FetchUser(username)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch username"})
		return
	}

	if user == (userModel.FetchUserResult{}) {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Username is unavailable"})
}