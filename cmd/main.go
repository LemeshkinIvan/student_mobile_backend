package main

import (
	"app/internal/bootstrap"
	"log"
)

func main() {
	app := bootstrap.InitApp()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
