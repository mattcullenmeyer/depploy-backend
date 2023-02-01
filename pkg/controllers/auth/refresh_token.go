package authController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type RefreshTokenPayload struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func RefreshToken(c *gin.Context) {
	var payload RefreshTokenPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken := payload.RefreshToken

	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is invalid"})
		return
	}

	fetchUserByAccountArgs := userModel.FetchUserByAccountParams{
		AccountId: claims.AccountId,
	}

	user, err := userModel.FetchUserByAccount(fetchUserByAccountArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Accounts can be blocked by setting verification status to false
	// therefore, check if account is blocked before refreshing an auth token
	if !user.Verified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You are not authorized to refresh your token"})
		return
	}

	generateTokenArgs := utils.GenerateTokenParams{
		Username:  claims.Username,
		AccountId: user.AccountId,
		Superuser: user.Superuser,
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh authentication token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth_token": token, "refresh_token": refresh})
}
