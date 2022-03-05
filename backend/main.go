package main

import (
	"github.com/Andi-App/Andi/database"
	router "github.com/Andi-App/Andi/routes"
)

func main() {
	database.Init()

	router.LoadRoutes()
}
