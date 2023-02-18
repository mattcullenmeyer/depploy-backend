package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubOauthTokenResult struct {
	AccessToken string `json:"access_token"`
}

type GitHubUserResult struct {
	AccountId int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
}

// https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/authorizing-oauth-apps
func GetGitHubOauthToken(code string) (string, error) {
	clientId, err := GetParameter("GitHubOauthClientId")
	if err != nil {
		return "", err
	}
	clientSecret, err := GetParameter("GitHubOauthClientSecret")
	if err != nil {
		return "", err
	}

	requestBody, _ := json.Marshal(map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"code":          code,
	})

	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	gitHubOauthTokenResult := GitHubOauthTokenResult{}
	if err := json.Unmarshal(respBody, &gitHubOauthTokenResult); err != nil {
		return "", err
	}

	return gitHubOauthTokenResult.AccessToken, nil
}

func GetGitHubUserData(token string) (GitHubUserResult, error) {
	emptyResult := GitHubUserResult{}

	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if err != nil {
		return emptyResult, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return emptyResult, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return emptyResult, err
	}

	// Check all the key values that are returned
	// fmt.Println("GitHub User Data: ", string(respBody))

	gitHubUserResult := GitHubUserResult{}
	if err := json.Unmarshal(respBody, &gitHubUserResult); err != nil {
		return emptyResult, err
	}

	return gitHubUserResult, nil
}
