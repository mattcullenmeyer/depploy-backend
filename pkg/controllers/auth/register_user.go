package authController

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
	"github.com/pquerna/otp/totp"
)

type RegisterUserPayload struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// https://pkg.go.dev/github.com/gin-gonic/gin#section-readme
// See "Model binding and validation" section
func RegisterUser(c *gin.Context) {
	var payload RegisterUserPayload

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

	createUserArgs := authModel.CreateUserParams{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	if err := authModel.CreateUser(createUserArgs); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save new user"})
		return
	}

	// Generate verification code
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "depploy.io",
		AccountName: "hello@depploy.io",
	})
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate verification code"})
		return
	}

	args := authModel.CreateVerificationCodeParams{
		Otp:      key.Secret(),
		Username: username,
		Email:    email,
	}

	// Save verification code to database
	if err := authModel.CreateVerificationCode(args); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save verification code"})
		return
	}

	// Send verification email

	c.JSON(http.StatusCreated, gin.H{"username": username, "otp": key.Secret()})
}
