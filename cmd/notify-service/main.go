package main

import (
	"context"
	"log"

	"notify-service/internal/app"

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
		log.Fatalf("error occured while creating app: %v", err)
	}

	if err := server.Start(context.Background()); err != nil {
		log.Fatalf("error occured while running server: %v", err)
	}
}
