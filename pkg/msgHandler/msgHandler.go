package msgHandler

import "context"

type HandleFunc func(ctx context.Context, msg []byte) error

type MsgHandler interface {
	ServeMSG(ctx context.Context, msg []byte) error
}

type handler struct {
	handler HandleFunc
}

func New(fn HandleFunc) MsgHandler {
	return &handler{handler: fn}
}

func (h *handler) ServeMSG(ctx context.Context, msg []byte) error {
	return h.handler(ctx, msg)
}
