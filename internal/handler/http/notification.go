package http

import (
	httpServError "Lists-app/internal/handler/error"
	msg "Lists-app/internal/handler/model/msg"
	"Lists-app/internal/handler/model/msg/event"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) AddNotify(c *gin.Context) {
	var notify msg.Test

	if err := c.BindJSON(&notify); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
	}

	ht := msg.Data{
		Data: []byte(notify.Data),
	}

	m := msg.MSG{
		Type:    event.AddNotify,
		Content: ht,
	}

	err := h.srv.Notification().Add(context.Background(), m)
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(http.StatusOK, "grge")

}

func (h *handler) GetCurrentNotify(c *gin.Context) {
	message, err := h.srv.Notification().GetCurrent(context.Background())
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(httpServError.Resolver(err), message)
}

func (h *handler) GetOldNotify(c *gin.Context) {
	message, err := h.srv.Notification().GetOld(context.Background())
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(httpServError.Resolver(err), message)
}
