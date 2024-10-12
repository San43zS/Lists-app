package rabbit

import (
	"Lists-app/internal/server/launcher"
	"context"
)

type server struct {
	srv launcher.Server
}

func New() launcher.Server {

	return nil
}

func (s server) Serve(ctx context.Context) error {
	return nil
}
