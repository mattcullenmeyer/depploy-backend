package authController

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
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

	verificationCode := payload.VerificationCode

	result, err := authModel.FetchVerificationCode(verificationCode)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}

	if result == (authModel.FetchVerificationCodeResult{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Verification code is invalid or expired"})
		return
	}

	if err := authModel.UpdateUserVerified(result.Username); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully verified email for %s", result.Email)})
}
