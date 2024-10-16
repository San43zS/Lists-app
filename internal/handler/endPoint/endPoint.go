package endPoint

import (
	"Lists-app/internal/service"
	"Lists-app/pkg/msgHandler"
	"context"
)

type handler struct {
	srv    service.Service
	router msgHandler.MsgHandler
}

func New(srv service.Service) msgHandler.MsgHandler {
	endPointParser := func(ctx context.Context, msg []byte) error {
		return nil
	}

	handler := &handler{
		srv:    srv,
		router: msgHandler.New(endPointParser),
	}

	return handler.router
}
