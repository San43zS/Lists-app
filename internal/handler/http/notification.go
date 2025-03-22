package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	httpServError "notify-service/internal/handler/error"
	event2 "notify-service/internal/handler/event"
	msg "notify-service/internal/handler/model/msg"
	"notify-service/internal/handler/model/msg/event"
	"time"
)

func (h handler) AddNotify(c *gin.Context) {
	var notify msg.Notify

	if err := c.BindJSON(&notify); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
	}

	m := msg.MSG{
		Type: event.AddNotify,
		Data: []byte(notify.Data),
		TTL:  notify.TTL,
	}

	err := h.srv.Notification().Add(context.Background(), m)
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(http.StatusOK, "Notification successfully added")
}

func (h handler) GetCurrentNotify(c *gin.Context) {

	m := msg.MSG{
		Type: event.GetCurrentNotify,
	}

	err := h.srv.Notification().Add(context.Background(), m)
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	time.Sleep(2000 * time.Millisecond)
	c.JSONP(http.StatusOK, event2.MasCurrent)
}

func (h handler) GetOldNotify(c *gin.Context) {

	m := msg.MSG{
		Type: event.GetOldNotify,
	}

	err := h.srv.Notification().Add(context.Background(), m)
	if err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	time.Sleep(2000 * time.Millisecond)
	c.JSONP(http.StatusOK, event2.MasOld)

}
