package handler

import (
	"Lists-app/internal/broker"
	"Lists-app/internal/handler/endPoint"
	"Lists-app/internal/handler/http"
	"Lists-app/internal/service"
	"Lists-app/pkg/msgHandler"
)

type Handler struct {
	EndPoint msgHandler.MsgHandler
	services service.Service
	Http     http.Handler
}

func New(services service.Service, broker broker.Broker) *Handler {
	return &Handler{
		EndPoint: endPoint.New(services),
		Http:     http.New(services),
	}
}
