package authController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type VerifyEmailPayload struct {
	VerificationCode string `json:"verification_code" binding:"required"`
}

func VerifyEmail(c *gin.Context) {
	var payload VerifyEmailPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deleteOtpArgs := authModel.DeleteOtpParams{
		Otp: payload.VerificationCode,
	}

	result, err := authModel.DeleteOtp(deleteOtpArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}

	if result == (authModel.DeleteOtpResult{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verification code is invalid or expired"})
		return
	}

	updateUserVerifiedParams := authModel.UpdateUserVerifiedParams{
		AccountId: result.AccountId,
		Verified:  true,
	}

	if err := authModel.UpdateUserVerified(updateUserVerifiedParams); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}

	generateTokenArgs := utils.GenerateTokenParams{
		AccountId:  result.AccountId,
		SuperAdmin: false,
	}

	authTokens, err := utils.GenerateAuthTokens(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate authentication tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth_token": authTokens.AuthToken, "refresh_token": authTokens.RefreshToken})
}
