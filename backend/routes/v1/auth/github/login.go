package github

import (
	"fmt"
	"net/http"

	"github.com/Andi-App/Andi/utils"
	"github.com/gin-gonic/gin"
)

func GithubLoginRedirect(c *gin.Context) {
	githubClientID := utils.GetGithubClientID()

	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		githubClientID,
		"http://localhost:4000/api/v1/aut/github/callback",
	)

	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}
