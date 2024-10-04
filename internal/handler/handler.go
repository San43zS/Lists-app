package handler

import (
	"Lists-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Service
}

func New(services service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Endpoints fo registration & authorization
	auth := router.Group("/user")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-out", h.signOut)
	}

	// api for working with todo lists & their items
	api := router.Group("/api")
	{
		api.POST("/", h.viewAllNotify)
		api.GET("/", h.createNotify)
		// for anyone who has connection with list on id
		api.GET("/:id", h.getListById)
		api.PUT("/:id", h.updateList)
		api.DELETE("/:id", h.deleteList)

	}
	return router
}
