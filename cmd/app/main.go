package main

import (
	"os"

	"github.com/Liar233/Task-tracker/internal/app"
	"github.com/spf13/viper"
)

func main() {

	appConfig, err := loadConfig()

	if err != nil {

		println(err.Error())

		os.Exit(2)

	}

	tasksApp := app.NewApp(appConfig)

	tasksApp.Bootstrap()

	if err = tasksApp.Run(); err != nil {

		println(err.Error())

		os.Exit(2)
	}
}

func loadConfig() (*app.ApplicationConfig, error) {

	vp := viper.New()
	vp.AutomaticEnv()

	config := &app.ApplicationConfig{}

	config.HttpPort = vp.GetUint64("HTTP_PORT")
	config.HttpHost = vp.GetString("HTTP_HOST")

	config.DBPort = vp.GetUint64("DB_PORT")
	config.DBHost = vp.GetString("DB_HOST")
	config.DBName = vp.GetString("DB_NAME")
	config.DBUser = vp.GetString("DB_USER")
	config.DBPassword = vp.GetString("DB_PASSWORD")

	return config, nil
}
