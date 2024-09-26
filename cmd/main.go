package main

import (
	Lists_app "Lists-app"
	"Lists-app/pkg/handlers"
	"Lists-app/pkg/repository"
	"Lists-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	server := new(Lists_app.Server)
	if err := server.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	//name of directory
	viper.AddConfigPath("configs")
	//name of file
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// go build -ldflags="-s -w"
