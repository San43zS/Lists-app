package main

import (
	"Lists-app/internal/handler"
	Lists_app "Lists-app/internal/server"
	"Lists-app/internal/service"
	"Lists-app/internal/storage"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	storages := storage.New()
	services := service.New(storages)
	handlers := handler.New(services)

	// TODO server := Lists_app.New()
	var server = Lists_app.New(":8080", handlers.InitRoutes())

	if err := server.Run(); err != nil {
		log.Fatalf("error occured while running http server: %v", err)
	}
}
