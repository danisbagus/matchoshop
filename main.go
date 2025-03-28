package main

import (
	"os"

	app "github.com/danisbagus/matchoshop/app/api"
)

func main() {
	// app.StartApp()
	e := app.Init()
	appPort := os.Getenv("PORT") // todo: move to config
	e.Logger.Fatal(e.Start(":" + appPort))
}
