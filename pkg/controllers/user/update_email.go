package userController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
)

type UpdateEmailPayload struct {
	Email string `json:"email" binding:"required"`
}

func UpdateEmail(c *gin.Context) {
	var payload UpdateEmailPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUserEmailArgs := userModel.UpdateUserEmailParams{
		AccountId: c.MustGet("accountId").(string),
		Email:     payload.Email,
	}

	if err := userModel.UpdateUserEmail(updateUserEmailArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated email successfully"})
}
