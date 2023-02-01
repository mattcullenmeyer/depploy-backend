package userController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

type UpdateUsernamePayload struct {
	Username string `json:"username" binding:"required"`
}

func UpdateUsername(c *gin.Context) {
	var payload UpdateUsernamePayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateAccountUsernameArgs := userModel.UpdateAccountUsernameParams{
		Username:  payload.Username,
		AccountId: c.MustGet("accountId").(string),
	}

	if err := userModel.UpdateAccountUsername(updateAccountUsernameArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated username"})
}
