package utils

import (
	"log"
	"os"
)

func GetGithubClientID() string {
	githubClientID, exists := os.LookupEnv("GITHUB_CLIENT_ID")
	if !exists {
		log.Fatal("Github Client ID not defined in .env file")
	}

	return githubClientID
}

func GetGithubClientSecret() string {
	githubClientSecret, exists := os.LookupEnv("GITHUB_CLIENT_SECRET")
	if !exists {
		log.Fatal("Github Client Secret not defined in .env file")
	}

	return githubClientSecret
}
