package main

import (
	// "app/factories"
	"app/packages/database"
	"app/routes"
)

func main() {
	// Run App
	database.Init()

	e := routes.InitV1()
	e.Logger.Fatal(e.Start(":1323"))

	// Run Seeders
	// factories.Factory()
}
