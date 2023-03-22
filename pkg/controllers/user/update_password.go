package userController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type UpdatePasswordParams struct {
	Password string `json:"password" binding:"required"`
}

func UpdatePassword(c *gin.Context) {
	var payload UpdatePasswordParams

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	updateUserPasswordArgs := userModel.UpdateUserPasswordParams{
		AccountId: c.MustGet("accountId").(string),
		Password:  hashedPassword,
	}

	if err := userModel.UpdateUserPassword(updateUserPasswordArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated password successfully"})
}
