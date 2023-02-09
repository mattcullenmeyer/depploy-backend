package authController

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
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

	googleUser, err := utils.GetGoogleUserData(token)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=google", redirectLocation))
		return
	}

	fetchUserByAccountArgs := userModel.FetchUserByAccountParams{
		AccountId: googleUser.AccountId,
	}

	user, err := userModel.FetchUserByAccount(fetchUserByAccountArgs)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=internal", redirectLocation))
		return
	}

	superAdmin := false
	isNewAccount := user == userModel.FetchUserByAccountResult{}

	if isNewAccount {
		createGoogleUserArgs := authModel.CreateGoogleUserParams{
			AccountId: googleUser.AccountId,
			Email:     googleUser.Email,
			Verified:  googleUser.Verified,
			Name:      googleUser.Name,
		}

		if err := authModel.CreateGoogleUser(createGoogleUserArgs); err != nil {
			log.Println(err.Error())
			c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=internal", redirectLocation))
			return
		}
	} else {
		superAdmin = user.SuperAdmin
	}

	generateTokenArgs := utils.GenerateTokenParams{
		AccountId:  googleUser.AccountId,
		SuperAdmin: superAdmin,
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

	if isNewAccount {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup/username", redirectLocation))
	} else {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/%s", redirectLocation, user.Username))
	}
}
