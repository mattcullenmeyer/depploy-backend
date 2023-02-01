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
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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

	username, email, password := payload.Username, payload.Email, payload.Password

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password encryption failed"})
		return
	}

	accountId := ksuid.New()

	createEmailUserArgs := authModel.CreateEmailUserParams{
		AccountId: accountId.String(),
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
	}

	if err := authModel.CreateEmailUser(createEmailUserArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
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
		Email:    email,
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
		Email:    email,
	}

	// Send verification email
	if err := utils.SendConfirmationEmail(emailArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email verification"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"username": username, "email": email})
}
