package main

import (
	"log"

	"Lists-app/internal/app"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	server, err := app.New()
	if err != nil {
		log.Fatalf("error occured while creating server: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("error occured while running http server: %v", err)
	}
}
