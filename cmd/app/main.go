package main

import (
	"os"

	"github.com/Liar233/Task-tracker/internal/app"
)

func main() {

	appConfig := app.ApplicationConfig{}

	tasksApp := app.NewApp(&appConfig)

	if err := tasksApp.Bootstrap(); err != nil {

		os.Exit(1)
	}

	if err := tasksApp.Run(); err != nil {

		os.Exit(2)
	}
}
