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
	Email string `json:"email" binding:"required"`
}

func ResendEmail(c *gin.Context) {
	var payload ResendEmailPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := payload.Email

	fetchUserByEmailArgs := userModel.FetchUserByEmailParams{
		Email: email,
	}

	user, err := userModel.FetchUserByEmail(fetchUserByEmailArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	if user == (userModel.FetchUserByEmailResult{}) {
		// User does not exist
		// Return status ok since we don't want to communicate that the user doesn't exist
		log.Printf("Cannot resend email verification because '%s' does not exist", email)
		c.Status(http.StatusOK)
		return
	}

	if user.Verified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is already verified"})
		return
	}

	generateOtpArgs := utils.GenerateOtpParams{
		Email: email,
	}

	otp, err := utils.GenerateOtp(generateOtpArgs)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate one-time password"})
		return
	}

	createOtpArgs := authModel.CreateOtpParams{
		Otp:       otp,
		AccountId: user.AccountId,
		Email:     user.Email,
	}

	// Save verification code to database
	if err := authModel.CreateOtp(createOtpArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save one-time password"})
		return
	}

	emailArgs := utils.SendConfirmationEmailParams{
		Otp:   otp,
		Email: user.Email,
	}

	if err := utils.SendConfirmationEmail(emailArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return
	}

	c.Status(http.StatusNoContent)
	// c.JSON(http.StatusOK, gin.H{"otp": otp})
}
