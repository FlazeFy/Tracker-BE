package main

import (
	"tracker/packages/database"
	"tracker/routes"
)

func main() {
	database.Init()
	e := routes.InitV1()

	e.Logger.Fatal(e.Start(":1323"))
}
