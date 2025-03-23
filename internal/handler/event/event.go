package event

import (
	"notify-service/internal/broker/rabbit/consumer"
	"notify-service/internal/broker/rabbit/producer"
	msg2 "notify-service/internal/handler/model/msg"
	"notify-service/internal/handler/model/msg/event"
	"notify-service/internal/service"
	"notify-service/pkg/msgHandler"
)

type handler struct {
	srv    service.Service
	router msgHandler.MsgHandler

	respondConsumer respCons
}

type respCons struct {
	p producer.Producer
	c consumer.Consumer
}

func New(srv service.Service) msgHandler.MsgHandler {
	endPointParser := func(msg msg2.MSG) (string, error) {
		return msg.Type, nil
	}

	handler := &handler{
		srv:    srv,
		router: msgHandler.New(endPointParser),
	}

	handler.initHandler()

	return handler.router
}

func (h handler) initHandler() {
	h.router.Add(event.AddNotify, h.AddNotify)
	h.router.Add(event.GetCurrentNotify, h.GetCurrentNotify)
	h.router.Add(event.GetOldNotify, h.GetOldNotify)
}
