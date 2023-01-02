package authController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	generateTokenArgs := utils.GenerateTokenParams{
		Username: claims.Username,
		Account:  claims.Account,
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

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refresh})
}
