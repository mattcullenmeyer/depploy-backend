package authController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type LoginUserPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(c *gin.Context) {
	var payload LoginUserPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := payload.Username

	user, err := userModel.FetchUser(username)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log in user"})
		return
	}

	if user == (userModel.FetchUserResult{}) {
		// User does not exist
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	if !user.Verified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please verify your email"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		// Password is incorrect
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	generateTokenArgs := utils.GenerateTokenParams{
		Username: username,
		Account:  username,
	}

	token, err := utils.GenerateToken(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	refresh, err := utils.GenerateRefreshToken(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refresh})
}
