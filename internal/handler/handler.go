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
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-out", h.signOut)
	}

	api := router.Group("/api")
	{
		api.POST("/test", h.test)
		api.POST("/", h.viewAllNotify)
		api.PUT("/:id", h.createNotify)
		api.DELETE("/:id", h.deleteNotify)
	}
	return router
}
