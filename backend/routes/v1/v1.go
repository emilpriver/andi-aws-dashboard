package v1

import (
	v1Auth "github.com/emilpriver/andi-aws-dashboard/routes/v1/auth"
	"github.com/gin-gonic/gin"
)

func V1Routes(router *gin.RouterGroup) {
	r := router.Group("/v1")
	{
		v1Auth.AuthRoutes(r)
	}
}
