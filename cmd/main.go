package main

import (
	Lists_app "Lists-app"
	"Lists-app/pkg/handlers"
	"log"
)

func main() {
	handlers := new(handlers.Handler)

	server := new(Lists_app.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

// go build -ldflags="-s -w"
