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

	// Fetch verification code and return username
	result, err := authModel.FetchVerificationCode(verificationCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// Update user as verified
	if err := authModel.UpdateUserVerification(result.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Email verified successfully for %s", result.Email)})
}
