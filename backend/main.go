package main

import (
	"github.com/emilpriver/andi-aws-dashboard/database"
	router "github.com/emilpriver/andi-aws-dashboard/routes"
	"github.com/emilpriver/andi-aws-dashboard/utils"
)

func main() {
	utils.LoadSecrets()

	database.Init()

	router.LoadRoutes()
}
