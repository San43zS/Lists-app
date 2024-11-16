package handler

import (
	"Lists-app/internal/broker"
	"Lists-app/internal/handler/event"
	"Lists-app/internal/handler/http"
	"Lists-app/internal/service"
	"Lists-app/pkg/msgHandler"
	http2 "net/http"
)

type Handler struct {
	EndPoint msgHandler.MsgHandler
	services service.Service
	Http     http2.Handler
}

func New(services service.Service, broker broker.Broker) *Handler {
	router := event.New(services)
	return &Handler{
		services: services,
		EndPoint: router,
		Http:     http.New(services),
	}
}
