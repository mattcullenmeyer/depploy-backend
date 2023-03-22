package userController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func CheckEmailAvailability(c *gin.Context) {
	email := c.Params.ByName("email")

	fetchUserByEmailArgs := userModel.FetchUserByEmailParams{
		Email: email,
	}

	user, err := userModel.FetchUserByEmail(fetchUserByEmailArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email"})
		return
	}

	if user == (userModel.FetchUserByEmailResult{}) {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email is unavailable"})
}
