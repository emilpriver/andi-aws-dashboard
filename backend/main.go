package main

import (
	"github.com/Andi-App/Andi/database"
	router "github.com/Andi-App/Andi/routes"
	"github.com/Andi-App/Andi/utils"
)

func main() {
	utils.LoadSecrets()

	database.Init()

	router.LoadRoutes()
}
