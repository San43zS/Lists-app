package app

import (
	"Lists-app/internal/handler"
	"Lists-app/internal/server/http"
	"Lists-app/internal/service"
	"Lists-app/internal/storage/config"
	"Lists-app/internal/storage/db/psql"
	"context"
)

type App struct {
	server *http.Server
}

func New() (App, error) {
	test := config.NewConfig()

	storage, err := psql.New(test)
	if err != nil {
		return App{}, err
	}

	srv := service.New(storage)

	handlers := handler.New(srv)
	app := App{
		server: http.New(handlers.InitRoutes()),
	}
	return app, nil
}

func (a *App) Run() error {
	return a.server.Run()
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
