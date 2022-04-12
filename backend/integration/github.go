package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/emilpriver/andi-aws-dashboard/types"
	"github.com/emilpriver/andi-aws-dashboard/utils"
)

func GithubGetUserAccessToken(callbackCode string) (string, error) {
	githubClientID := utils.GetGithubClientID()
	githubClientSecret := utils.GetGithubClientSecret()
	githubClientRediretUri := utils.GetGithubClientRedirectUri()

	values := map[string]string{
		"client_id":     githubClientID,
		"client_secret": githubClientSecret,
		"redirect_uri":  githubClientRediretUri,
		"code":          callbackCode,
	}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{}

	request, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(json_data),
	)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	if err != nil {
		return "", errors.New("Failed to get access token")
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
		return "", errors.New("Failed to send reqiest to Github")
	}

	defer response.Body.Close()

	var data types.GithubLoginCallbackResponse

	json.NewDecoder(response.Body).Decode(&data)

	if len(data.AccessToken) == 0 {
		return "", errors.New("Failed to get access token")
	}

	return data.AccessToken, nil
}

func GithubGetUserByToken(token string) (types.GithubUser, error) {
	client := http.Client{}

	request, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	authToken := fmt.Sprintf("token %v", token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", authToken)

	if err != nil {
		log.Fatal(err)

		return types.GithubUser{}, errors.New("Failed to build request")
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)

		return types.GithubUser{}, errors.New("Failed to send request to Github")
	}

	defer response.Body.Close()

	var data types.GithubUser

	json.NewDecoder(response.Body).Decode(&data)

	return data, nil
}
