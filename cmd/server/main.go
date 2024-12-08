package main

import (
	"tracker/app"
)

func main() {
	app.InitLogger()
	app.InitDb()

	server := app.NewServer()
	server.InitServer()
}
