package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	httpServError "notify-service/internal/handler/error"
	USER "notify-service/internal/model/user"
)

func (h handler) SignIn(c *gin.Context) {
	var user USER.User

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

func (h handler) SignUp(c *gin.Context) {
	var user USER.User

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

func (h handler) SignOut(c *gin.Context) {
	var user USER.User

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
