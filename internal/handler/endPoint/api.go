package endPoint

import (
	httpServError "Lists-app/internal/handler/error"
	"Lists-app/internal/model/notification"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) viewAllNotify(c *gin.Context) {

}

func (h *handler) createNotify(c *gin.Context) {

}

func (h *handler) deleteNotify(c *gin.Context) {

}

func (h *handler) test(c *gin.Context) {
	var notify notification.Notification

	if err := c.BindJSON(&notify); err != nil {

		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	if err := h.srv.Notification().Add(context.Background(), notify); err != nil {

		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, "You have successfully received a notification")

}
