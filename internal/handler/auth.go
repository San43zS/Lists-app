package handler

import (
	user2 "Lists-app/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Autorization.Registration; err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, gin.H{"message": "user created"})

}

func (h *Handler) signIn(c *gin.Context) {
	var user user2.User

	if err := c.BindJSON(&user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Authenticate(user); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, gin.H{"message": "user authorized"})
}

func (h *Handler) signOut(c *gin.Context) {

}
