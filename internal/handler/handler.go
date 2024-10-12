package handler

import (
	"Lists-app/internal/service"
	"Lists-app/pkg/msgHandler"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	EndPoint msgHandler.MsgHandler
	services service.Service
}

func New(services service.Service) *Handler {
	return &Handler{
		services: services,
		EndPoint: msgHandler.,
	}
}

type MSG struct {
	endpoint string
	body     []byte
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
