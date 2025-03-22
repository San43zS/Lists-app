package handler

import (
	HTTP "net/http"
	"notify-service/internal/broker"
	"notify-service/internal/handler/event"
	"notify-service/internal/handler/http"
	"notify-service/internal/service"
	"notify-service/pkg/msgHandler"
)

type Handler struct {
	EndPoint msgHandler.MsgHandler
	services service.Service
	Http     HTTP.Handler
}

func New(services service.Service, broker broker.Broker) *Handler {
	router := event.New(services)
	return &Handler{
		services: services,
		EndPoint: router,
		Http:     http.New(services),
	}
}
