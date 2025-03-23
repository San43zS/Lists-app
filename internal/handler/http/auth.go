package http

import (
	httpServError "notify-service/internal/handler/error"
	user2 "notify-service/internal/model/user"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) SignIn(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	if err := h.srv.User().SignIn(context.Background(), user); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, "You have successfully logged in to your account")
}

func (h *handler) SignUp(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	if err := h.srv.User().SignUp(context.Background(), user); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(http.StatusOK, "User successfully registered")
}

func (h *handler) SignOut(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	if err := h.srv.User().Delete(context.Background(), user); err != nil {
		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	c.JSONP(http.StatusOK, "user authorized")
}
