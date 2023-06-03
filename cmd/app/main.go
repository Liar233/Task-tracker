package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/Liar233/Task-tracker/internal/app"
)

func main() {

	appConfig, err := loadConfig()

	if err != nil {

		println(err.Error())

		os.Exit(2)

	}

	tasksApp := app.NewApp(appConfig)

	tasksApp.Bootstrap()

	if err := tasksApp.Run(); err != nil {

		println(err.Error())

		os.Exit(2)
	}
}

func loadConfig() (*app.ApplicationConfig, error) {

	envHost, ok := os.LookupEnv("HTTP_HOST")

	if !ok {

		return nil, errors.New("variable HTTP_HOST does not exists")
	}

	if envHost == "" {

		return nil, errors.New("variable HTTP_HOST not valid")
	}

	envPort, ok := os.LookupEnv("HTTP_PORT")

	if !ok {

		return nil, errors.New("variable HTTP_PORT does not exists")
	}

	port, err := strconv.ParseUint(envPort, 10, 64)

	if err != nil || port == uint64(0) {

		return nil, errors.New("variable HTTP_PORT does not valid")
	}

	return &app.ApplicationConfig{
		Host: envHost,
		Port: port,
	}, err
}
