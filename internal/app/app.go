package app

import (
	"Lists-app/internal/broker"
	"Lists-app/internal/server"
	"Lists-app/internal/service"
	"Lists-app/internal/storage/config"
	"Lists-app/internal/storage/db/psql"
	"context"
	"fmt"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("app")

type App struct {
	server service.Service
	broker broker.Broker
}

func New() (*App, error) {
	storage, err := psql.New(config.NewConfig())
	if err != nil {
		return &App{}, err
	}

	broker, err := broker.New()
	if err != nil {
		return &App{}, err
	}

	srv := service.New(storage, broker)

	app := &App{
		server: srv,
		broker: broker,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	srv, err := server.New(a.server, a.broker)
	if err != nil {
		return err
	}

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("server stopped with error: %w\n", err)
	}

	log.Infof("server stopped")
	return nil
}
