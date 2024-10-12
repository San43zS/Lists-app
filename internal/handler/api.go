package handler

import (
	httpServError "Lists-app/internal/handler/model/error"
	"Lists-app/internal/model/notification"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) viewAllNotify(c *gin.Context) {

}

func (h *Handler) createNotify(c *gin.Context) {

}

func (h *Handler) deleteNotify(c *gin.Context) {

}

func (h *Handler) test(c *gin.Context) {
	var notify notification.Notification

	if err := c.BindJSON(&notify); err != nil {

		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	if err := h.services.Notification().Add(context.Background(), notify); err != nil {

		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, "You have successfully received a notification")

}
