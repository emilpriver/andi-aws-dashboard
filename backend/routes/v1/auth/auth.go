package v1Auth

import (
	"github.com/emilpriver/andi-aws-dashboard/routes/v1/auth/github"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {
	r := router.Group("/auth")
	{
		githubRoutes := r.Group("/github")
		{
			githubRoutes.GET("/login", github.GithubLoginRedirect)
			githubRoutes.GET("/callback", github.GithubLoginCallback)
		}
	}
}
