package github

import (
	"fmt"
	"net/http"

	"github.com/Andi-App/Andi/utils"
	"github.com/gin-gonic/gin"
)

func GithubLoginRedirect(c *gin.Context) {
	githubClientID := utils.GetGithubClientID()
	githubRediretUri := utils.GetGithubClientRedirectUri()

	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user",
		githubClientID,
		githubRediretUri,
	)

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
