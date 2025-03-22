package event

import (
	"encoding/json"
	"fmt"
	HTTP "net/http"
	message "notify-service/internal/handler/model/msg"
	"notify-service/internal/handler/model/msg/event"
	"notify-service/internal/service"
	"notify-service/pkg/msgHandler"
)

type handler struct {
	srv     service.Service
	router  msgHandler.MsgHandler
	handler HTTP.Handler
}

func New(srv service.Service) msgHandler.MsgHandler {
	eventParser := func(msg []byte) (string, error) {
		var common message.STRUCT
		if err := json.Unmarshal(msg, &common); err != nil {
			return "", fmt.Errorf("error while parsing msg: %w", err)
		}
		return common.Type, nil
	}

	handler := &handler{
		srv:    srv,
		router: msgHandler.New(eventParser),
	}

	handler.initHandler()

	return handler.router
}

func (h handler) initHandler() {

	h.router.Add(event.ShowCurrent, h.ShowCurrent)
	h.router.Add(event.ShowOld, h.ShowOld)

}
