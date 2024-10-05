package handler

import (
	user2 "Lists-app/internal/model/user"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func (h *Handler) signUp(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(http.StatusBadRequest, "Error receiving data")
		return
	}

	if _, err := h.services.User().Verify(context.Background(), user); err != nil {
		c.JSONP(http.StatusBadRequest, "Error verifying user: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, "user already exists")

}

func (h *Handler) signIn(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.User().Insert(context.Background(), user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, gin.H{"message": "user authorized"})
}

func (h *Handler) signOut(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.User().Delete(context.Background(), user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, gin.H{"message": "user authorized"})
}
