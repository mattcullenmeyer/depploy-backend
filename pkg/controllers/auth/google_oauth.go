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

	accountId := fmt.Sprintf("GO%s", googleUser.AccountId) // Google account IDs are prefixed with "GO"

	fetchUserByAccountArgs := userModel.FetchUserByAccountParams{
		AccountId: accountId,
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
		createOauthUserArgs := authModel.CreateOauthUserParams{
			AccountId:          accountId,
			Email:              googleUser.Email,
			Name:               googleUser.Name,
			RegistrationMethod: "Google",
		}

		if err := authModel.CreateOauthUser(createOauthUserArgs); err != nil {
			log.Println(err.Error())
			c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=internal", redirectLocation))
			return
		}
	} else {
		superAdmin = user.SuperAdmin
	}

	generateTokenArgs := utils.GenerateTokenParams{
		AccountId:  accountId,
		SuperAdmin: superAdmin,
	}

	authTokens, err := utils.GenerateAuthTokens(generateTokenArgs)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=internal", redirectLocation))
		return
	}

	utils.SetAuthCookies(c, authTokens)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
