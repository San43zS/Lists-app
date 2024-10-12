package server

import (
	"context"

	"Lists-app/internal/broker"
	"Lists-app/internal/server/launcher"

	"Lists-app/internal/server/launcher/rabbit"

	"Lists-app/internal/server/launcher/http"

	"Lists-app/internal/service"
)

type server struct {
	servers []launcher.Server
}

func New(srv service.Service) (launcher.Server, error) {
	broker, err := broker.New()
	if err != nil {
		return nil, err
	}

	server := &server{
		servers: []launcher.Server{
			rabbit.New(broker.RabbitMQ, srv),
			http.New(),
		},
	}

	return server, nil
}

func (s *server) Serve(ctx context.Context) error {
	return nil
}

func (s *server) serve(ctx context.Context) error {
	return nil
}
