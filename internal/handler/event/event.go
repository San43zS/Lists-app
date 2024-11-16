package endPoint

import (
	"Lists-app/internal/service"
	"Lists-app/pkg/msgHandler"
	"context"
	"github.com/gin-gonic/gin"
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
	handler.initHandler()

	return handler.router
}

func (h handler) initHandler() *gin.Engine {

	h.router.Add(context.Background(), h.AddNotify)
}
