package authController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
	"github.com/segmentio/ksuid"
)

type RegisterEmailUserPayload struct {
	Email string `json:"email" binding:"required"`
}

// https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
// See "Model binding and validation" section
func RegisterEmailUser(c *gin.Context) {
	var payload RegisterEmailUserPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountId := ksuid.New().String()
	email := payload.Email

	createEmailUserArgs := authModel.CreateEmailUserParams{
		AccountId: accountId,
		Email:     email,
	}

	if err := authModel.CreateEmailUser(createEmailUserArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered"})
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
		AccountId: accountId,
		Email:     email,
	}

	// Save one-time password to database
	if err := authModel.CreateOtp(createOtpArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save one-time password"})
		return
	}

	emailArgs := utils.SendConfirmationEmailParams{
		Otp:   otp,
		Email: email,
	}

	// Send verification email
	if err := utils.SendConfirmationEmail(emailArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return
	}

	// TODO: Redact part of email
	c.JSON(http.StatusCreated, gin.H{"email": email})
}
