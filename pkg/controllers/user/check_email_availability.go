package userController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

func CheckEmailAvailability(c *gin.Context) {
	email := c.Params.ByName("email")

	user, err := userModel.FetchUser(email)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch email address"})
		return
	}

	if user == (userModel.FetchUserResult{}) {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email address is unavailable"})
}
