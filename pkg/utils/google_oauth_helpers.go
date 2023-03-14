package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserResult struct {
	AccountId string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
}

func GetOauthConfig() (*oauth2.Config, error) {
	clientId, err := GetParameter("GoogleOauthClientId")
	if err != nil {
		return &oauth2.Config{}, err
	}
	clientSecret, err := GetParameter("GoogleOauthClientSecret")
	if err != nil {
		return &oauth2.Config{}, err
	}

	backendHost := os.Getenv("BACKEND_HOST")
	redirectUrl := fmt.Sprintf("%s/auth/google", backendHost)

	conf := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}

	return conf, nil
}

// https://pkg.go.dev/golang.org/x/oauth2#example-Config
func GetGoogleOauthToken(code string) (*oauth2.Token, error) {
	conf, _ := GetOauthConfig()

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return token, err
	}

	return token, nil
}

func GetGoogleUserData(token *oauth2.Token) (GoogleUserResult, error) {
	conf, _ := GetOauthConfig()

	client := conf.Client(context.Background(), token)

	emptyResult := GoogleUserResult{}

	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token=%s", token.AccessToken)

	resp, err := client.Get(url)
	if err != nil {
		return emptyResult, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyResult, err
	}

	// Check all the key values that are returned
	// fmt.Println("Google User Data: ", string(respBody))

	googleUserResult := GoogleUserResult{}
	if err := json.Unmarshal(respBody, &googleUserResult); err != nil {
		return emptyResult, err
	}

	return googleUserResult, nil
}
