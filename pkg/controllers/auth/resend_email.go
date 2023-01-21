package authController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

type ResendEmailPayload struct {
	Username string `json:"username" binding:"required"`
}

func ResendEmail(c *gin.Context) {
	var payload ResendEmailPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := payload.Username

	user, err := userModel.FetchUser(username)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == (userModel.FetchUserResult{}) {
		// User does not exist
		// Return status ok since we don't want to communicate that the user doesn't exist
		log.Printf("Cannot resend email verification because '%s' does not exist", username)
		c.Status(http.StatusOK)
		return
	}

	if user.Verified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is already verified"})
		return
	}

	otp, err := utils.GenerateOtp()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate one-time password"})
		return
	}

	otpArgs := authModel.CreateVerificationCodeParams{
		Otp:      otp,
		Username: username,
		Email:    user.Email,
	}

	// Save verification code to database
	if err := authModel.CreateVerificationCode(otpArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save one-time password"})
		return
	}

	emailArgs := utils.SendConfirmationEmailParams{
		Otp:      otp,
		Username: username,
		Email:    user.Email,
	}

	if err := utils.SendConfirmationEmail(emailArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return
	}

	c.Status(http.StatusNoContent)
	// c.JSON(http.StatusOK, gin.H{"otp": otp})
}
