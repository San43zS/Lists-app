package msgHandler

import (
	msg2 "Lists-app/internal/handler/model/msg"
	"context"
	"errors"
)

type EventParser func(msg msg2.MSG) (string, error)
type HandlerFunc1 func(ctx context.Context, msg msg2.MSG) error
type HandlerFunc2 func(ctx context.Context) ([]byte, error)

type MsgResolver interface {
	ServeMSG(ctx context.Context, msg msg2.MSG) ([]byte, error)
}

type MsgHandler interface {
	MsgResolver
	Add(event string, fn interface{})
}

type handler struct {
	eventParser EventParser
	handlers    map[string]interface{}
}

func New(parser EventParser) MsgHandler {
	return &handler{
		eventParser: parser,
		handlers:    make(map[string]interface{}),
	}
}

func (h *handler) ServeMSG(ctx context.Context, msg msg2.MSG) ([]byte, error) {
	event, err := h.eventParser(msg)
	if err != nil {
		return nil, err
	}

	fn, ok := h.handlers[event]
	if !ok {
		return nil, nil
	}

	switch fn := fn.(type) {
	case HandlerFunc1:
		return nil, fn(ctx, msg)
	case HandlerFunc2:
		m, err := fn(ctx)
		return m, err
	default:
		return nil, errors.New("unknown handler type")
	}
}

func (h *handler) Add(event string, fn interface{}) {
	h.handlers[event] = fn
}
