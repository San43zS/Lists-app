package endPoint

import (
	"Lists-app/internal/service"
	"Lists-app/pkg/msgHandler"
)

type handler struct {
	srv    service.Service
	router msgHandler.MsgHandler
}

func New(srv service.Service) msgHandler.MsgHandler {
	handler := &handler{
		srv:    srv,
		router: msgHandler.New(srv),
	}

	return handler.router

}
