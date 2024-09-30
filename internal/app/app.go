package app

import (
	"Lists-app/internal/handler"
	"Lists-app/internal/server"
	"Lists-app/internal/service"
	"Lists-app/internal/storage/config"
	"Lists-app/internal/storage/db/psql"
	"context"
)

type app struct {
	server *server.Server
}

func New() (*app, error) {
	storage, err := psql.New(config.NewConfig())
	if err != nil {
		return nil, err
	}

	services := service.New(storage)
	handlers := handler.New(services)
	app := &app{
		server: server.New(handlers.InitRoutes()),
	}
	return app, nil
}

func (a *app) Run() error {
	return a.server.Run()
}

func (a *app) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
