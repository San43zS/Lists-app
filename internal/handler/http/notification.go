package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	httpServError "notify-service/internal/handler/error"
	msg "notify-service/internal/handler/model/msg"
	"notify-service/internal/handler/model/msg/event"
)

func (h handler) AddNotify(c *gin.Context) {
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

func (h handler) GetCurrentNotify(c *gin.Context) {
	message, err := h.srv.Notification().GetCurrent(context.Background())
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(httpServError.Resolver(err), message)
}

func (h handler) GetOldNotify(c *gin.Context) {
	message, err := h.srv.Notification().GetOld(context.Background())
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(httpServError.Resolver(err), message)
}
