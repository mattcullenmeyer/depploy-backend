package authController

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

func GoogleOAuth(c *gin.Context) {
	code := c.Query("code")
	// Also returns a "scope" query param

	redirectLocation := os.Getenv("CONSOLE_HOST")

	token, err := utils.GetGoogleOauthToken(code)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=google", redirectLocation))
		return
	}

	user, err := utils.GetGoogleUserData(token)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=google", redirectLocation))
		return
	}

	createGoogleUserArgs := authModel.CreateGoogleUserParams{
		AccountId: user.AccountId,
		Email:     user.Email,
		Verified:  user.Verified,
		Name:      user.Name,
	}

	if err := authModel.CreateGoogleUser(createGoogleUserArgs); err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=internal", redirectLocation))
		return
	}

	generateTokenArgs := utils.GenerateTokenParams{
		AccountId: user.AccountId,
		Superuser: false,
	}

	authToken, err := utils.GenerateToken(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=internal", redirectLocation))
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=internal", redirectLocation))
		return
	}

	in15Minutes := 15 * 60
	in24Hours := 24 * 60 * 60

	cookieDomain := os.Getenv("COOKIE_DOMAIN")

	c.SetCookie("auth_token", authToken, in15Minutes, "/", cookieDomain, true, false)
	c.SetCookie("refresh_token", refreshToken, in24Hours, "/", cookieDomain, true, false)

	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup/username", redirectLocation))
}
