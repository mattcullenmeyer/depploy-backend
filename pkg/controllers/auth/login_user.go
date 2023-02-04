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

	fetchUserByUsernameArgs := userModel.FetchUserByUsernameParams{
		Username: username,
	}

	user, err := userModel.FetchUserByUsername(fetchUserByUsernameArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log in user"})
		return
	}

	// Return 400 status code
	if user == (userModel.FetchUserByUsernameResult{}) {
		// User does not exist
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	// Return 400 status code
	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		// Password is incorrect
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	// Return 403 status code
	if !user.Verified {
		c.JSON(http.StatusForbidden, gin.H{"error": "Please verify your email address"})
		return
	}

	generateTokenArgs := utils.GenerateTokenParams{
		AccountId: user.AccountId,
		Superuser: user.Superuser,
	}

	authToken, err := utils.GenerateToken(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth_token": authToken, "refresh_token": refreshToken})
}
