package authController

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	authModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/auth"
	userModel "github.com/mattcullenmeyer/depploy-backend/pkg/models/user"
	"github.com/mattcullenmeyer/depploy-backend/pkg/utils"
)

func GitHubOAuth(c *gin.Context) {
	code := c.Query("code")
	// Also returns a "state" query param

	redirectLocation := os.Getenv("CONSOLE_HOST")

	token, err := utils.GetGitHubOauthToken(code)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=github", redirectLocation))
		return
	}

	gitHubUser, err := utils.GetGitHubUserData(token)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/signup?error=github", redirectLocation))
		return
	}

	gitHubAccountId := strconv.Itoa(gitHubUser.AccountId)
	accountId := fmt.Sprintf("GH%s", gitHubAccountId) // GitHub account IDs are prefixed with "GH"

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
			Email:              gitHubUser.Email,
			Name:               gitHubUser.Name,
			RegistrationMethod: "GitHub",
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
