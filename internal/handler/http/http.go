package http

import (
	"notify-service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	srv service.Service
}

func New(service service.Service) http.Handler {
	handler := handler{
		srv: service,
	}

	return handler.InitRoutes()
}

func (h handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-out", h.SignOut)
	}

	ws := router.Group("/ws")
	{
		ws.POST("/add", h.AddNotify)
		ws.GET("/get-current-notify", h.GetCurrentNotify)
		ws.GET("/get-old-notify", h.GetOldNotify)
	}

	return router
}
