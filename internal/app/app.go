package app

import (
	"context"
	"fmt"
	"github.com/op/go-logging"
	"notify-service/internal/broker"
	"notify-service/internal/server"
	"notify-service/internal/service"
	"notify-service/internal/storage/config"
	"notify-service/internal/storage/db/psql"
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

	//broker.RabbitMQ.Producer().Produce(context.Background(), []byte("test"))
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
