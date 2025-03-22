package server

import (
	"context"
	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"
	"notify-service/internal/handler"
	"sync"

	"notify-service/internal/broker"
	"notify-service/internal/server/launcher"

	"notify-service/internal/server/launcher/rabbit"

	"notify-service/internal/server/launcher/http"

	"notify-service/internal/service"
)

var log = logging.MustGetLogger("server")

type server struct {
	servers []launcher.Server
}

func New(srv service.Service, broker broker.Broker) (launcher.Server, error) {
	h := handler.New(srv, broker)

	server := &server{
		servers: []launcher.Server{
			http.New(h.Http),
			rabbit.New(broker.RabbitMQ, h.EndPoint),
		},
	}

	return server, nil
}

func (s *server) Serve(ctx context.Context) error {
	gr, grCtx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		return s.serve(grCtx)
	})

	var err error

	if err = gr.Wait(); err != nil {
		log.Criticalf("server stopped with error: %v", err)
	}

	log.Infof("app: shutting down the server...")

	return err
}

func (s *server) serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	gr, grCtx := errgroup.WithContext(ctx)

	for _, s := range s.servers {
		s := s

		gr.Go(func() error {
			defer wg.Done()

			return s.Serve(grCtx)
		})
	}

	wg.Wait()

	return gr.Wait()
}
