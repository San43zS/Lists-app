package http

import (
	"Lists-app/internal/server/launcher"
	"context"
	"errors"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type server struct {
	srv *http.Server
}

func New(handler http.Handler) launcher.Server {
	var httpServer = &http.Server{
		Addr:    "localhost" + ":" + viper.GetString("http_port"),
		Handler: handler,
		// Size of the HTTP request header allowed.
		MaxHeaderBytes: 1 << 20, // 1mb
		// Timeouts when server cannot read/write data from client for 10s
		// Then connection will be closed
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	srv := &server{
		srv: httpServer,
	}

	return srv
}

func (s *server) Serve(ctx context.Context) error {
	errCh := make(chan error)

	go func() {
		defer close(errCh)

		if err := s.srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():

	case err := <-errCh:
		return err
	}

	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}

	// TODO: modify graceful shutdown with error return

	return nil
}
