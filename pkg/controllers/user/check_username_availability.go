package userController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func CheckUsernameAvailability(c *gin.Context) {
	username := c.Params.ByName("username")

	fetchUserByUsernameArgs := userModel.FetchUserByUsernameParams{
		Username: username,
	}

	user, err := userModel.FetchUserByUsername(fetchUserByUsernameArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch username"})
		return
	}

	if user == (userModel.FetchUserByUsernameResult{}) {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Username is unavailable"})
}
