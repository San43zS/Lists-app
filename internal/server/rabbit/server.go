package rabbit

import (
	"Lists-app/internal/broker"
	"Lists-app/internal/server/launcher"
	"Lists-app/internal/server/launcher/rabbit"
	"Lists-app/internal/service"
	"context"
)

type server struct {
	servers []launcher.Server
}

func New(srv service.Service) launcher.Server {
	broker, err := broker.New()
	if err != nil {
		return nil
	}

	server := &server{
		servers: []launcher.Server{
			rabbit.New(),
		},
	}

	return server

}

func (s *server) Serve(ctx context.Context) error {
	return nil
}

func (s *server) serve(ctx context.Context) error {

}
