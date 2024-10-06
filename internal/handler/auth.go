package handler

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

func (h *Handler) signIn(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {

		c.JSONP(httpServError.ErrorResolver(err), err.Error())
		return
	}

	if err := h.services.User().Verify(context.Background(), user); err != nil {

		c.JSONP(httpServError.ErrorResolver(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, "You have successfully logged in to your account")

}

func (h *Handler) signUp(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(httpServError.ErrorResolver(err), err.Error())
		return
	}

	if err := h.services.User().Insert(context.Background(), user); err != nil {

		c.JSONP(httpServError.ErrorResolver(err), err.Error())
		return
	}

	c.JSONP(http.StatusOK, "User successfully registered")
}

func (h *Handler) signOut(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(httpServError.ErrorResolver(err), err.Error())
		return
	}

	if err := h.services.User().Delete(context.Background(), user); err != nil {
		c.JSONP(httpServError.ErrorResolver(err), err.Error())
		return
	}

	c.JSONP(http.StatusOK, "user authorized")
}
