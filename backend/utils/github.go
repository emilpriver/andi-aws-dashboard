package utils

import (
	"log"
)

func GetGithubClientID() string {
	githubClientID := EnvVariable("GITHUB_CLIENT_ID")

	if len(githubClientID) > 0 {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return githubClientID
}

func GetGithubClientSecret() string {
	githubClientSecret := EnvVariable("GITHUB_CLIENT_SECRET")

	if len(githubClientSecret) > 0 {
		log.Fatal("Github Client Secret not defined in .env file")
	}

	return githubClientSecret
}
