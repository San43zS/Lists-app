package http

import (
	httpServError "Lists-app/internal/handler/error"
	user2 "Lists-app/internal/model/user"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func (h *handler) SignIn(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {

		c.JSONP(httpServError.Resolver(err), err.Error())
		return
	}

	if err := h.srv.User().Verify(context.Background(), user); err != nil {

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

	if err := h.srv.User().Insert(context.Background(), user); err != nil {

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
