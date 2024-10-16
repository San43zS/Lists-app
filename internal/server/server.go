package server

import (
	"Lists-app/internal/handler"
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"sync"

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
	bkr, err := broker.New()
	if err != nil {

		return nil, err
	}

	h := handler.New(srv, bkr)
	hHttp := h.Http

	server := &server{
		servers: []launcher.Server{
			rabbit.New(bkr.RabbitMQ, h.EndPoint),
			http.New(hHttp.InitRoutes()),
		},
	}

	return server, nil
}

func (s *server) Serve(ctx context.Context) error {
	gr, grCtx := errgroup.WithContext(ctx)

	// start server
	gr.Go(func() error {
		return s.serve(grCtx)
	})

	var err error

	if err = gr.Wait(); err != nil {
		log.Println("server stopped with error: ", err)
	}

	log.Println("app: shutting down the server")

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
