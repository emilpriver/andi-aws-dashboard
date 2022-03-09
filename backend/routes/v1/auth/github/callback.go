package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Andi-App/Andi/utils"
	"github.com/gin-gonic/gin"
)

type GithubLoginCallbackResponse struct {
	AccessToken string `json:"access_token"`
}

type GithubUser struct {
	login                     string
	id                        int64
	node_id                   string
	avatar_url                string
	gravatar_id               string
	url                       string
	html_url                  string
	followers_url             string
	following_url             string
	gists_url                 string
	starred_url               string
	subscriptions_url         string
	organizations_url         string
	repos_url                 string
	events_url                string
	received_events_url       string
	site_admin                bool
	name                      string
	company                   string
	blog                      string
	location                  string
	email                     string
	hireable                  string
	bio                       string
	twitter_username          string
	public_repos              int64
	public_gists              int64
	followers                 int64
	following                 int64
	created_at                string
	updated_at                string
	private_gists             int64
	total_private_repos       int64
	owned_private_repos       int64
	disk_usage                int64
	collaborators             int64
	two_factor_authentication bool
}

func getUserAccessToken(callbackCode string) (string, bool) {
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
		log.Fatal(err)
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var data GithubLoginCallbackResponse

	json.NewDecoder(response.Body).Decode(&data)

	if len(data.AccessToken) == 0 {
		return "", false
	}

	return data.AccessToken, true
}

func getUserByToken(token string) GithubUser {
	client := http.Client{}

	request, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var data GithubUser

	json.NewDecoder(response.Body).Decode(&data)
	fmt.Println(data)
	return data
}

func GithubLoginCallback(c *gin.Context) {
	callbackCode, callbackCodeExists := c.GetQuery("code")

	if !callbackCodeExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Missing callback code",
		})
		return
	}

	accessToken, accessTokenExists := getUserAccessToken(callbackCode)

	if !accessTokenExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Coudn't get access token from Github",
		})

		return
	}

	user := getUserByToken(accessToken)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
