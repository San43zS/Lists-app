package app

import (
	"Lists-app/internal/server"
	"Lists-app/internal/service"
	"Lists-app/internal/storage/config"
	"Lists-app/internal/storage/db/psql"
	"context"
	"fmt"
	"log"
)

type App struct {
	server service.Service
}

func New() (*App, error) {
	test := config.NewConfig()

	storage, err := psql.New(test)
	if err != nil {

		return &App{}, err
	}

	srv := service.New(storage)

	app := &App{
		server: srv,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	srv, err := server.New(a.server)
	if err != nil {

		return err
	}

	if err := srv.Serve(ctx); err != nil {

		return fmt.Errorf("server stopped with error: %w\n", err)
	}

	log.Println("server stopped")
	return nil
}
