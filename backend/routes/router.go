package router

import (
	v1 "github.com/emilpriver/andi-aws-dashboard/routes/v1"
	"github.com/gin-gonic/gin"
)

func LoadRoutes() {
	r := gin.Default()

	api := r.Group("/api")
	{
		v1.V1Routes(api)
	}

	r.Run()
}
