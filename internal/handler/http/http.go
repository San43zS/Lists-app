package http

import (
	"Lists-app/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	InitRoutes() *gin.Engine
}

type handler struct {
	srv service.Service
}

func New(service service.Service) Handler {
	return &handler{
		srv: service,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Endpoints fo registration & authorization
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-out", h.SignOut)
	}

	return router
}
