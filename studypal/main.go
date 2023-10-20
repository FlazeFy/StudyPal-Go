package main

import (
	"studypal/packages/database"

	"studypal/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":3001"))
}
