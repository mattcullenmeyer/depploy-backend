package authController

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refreshToken := payload.RefreshToken

	claims, err := utils.ValidateToken(refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	generateTokenArgs := utils.GenerateTokenParams{
		Username: claims.Username,
		Account:  claims.Account,
	}

	// Generate JWT
	token, err := utils.GenerateToken(generateTokenArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Generate refresh JWT
	refresh, err := utils.GenerateRefreshToken(generateTokenArgs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refresh})
}
