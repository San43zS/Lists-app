package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"notify-service/internal/service"
)

type handler struct {
	srv      service.Service
	upgrader *websocket.Upgrader
}

func New(service service.Service) http.Handler {

	handler := &handler{
		srv: service,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	return handler.InitRoutes(handler.upgrader)
}

func (h handler) InitRoutes(up *websocket.Upgrader) *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-out", h.SignOut)
	}

	notify := router.Group("/notify")
	{
		notify.POST("/add", h.AddNotify)
		notify.POST("/get-current-notify", h.GetCurrentNotify)
		notify.GET("/get-old-notify", h.GetOldNotify)
	}

	ws := router.Group("/ws")
	{
		ws.Use(func(c *gin.Context) {
			conn, err := configuration(c, up)
			if err != nil {
				return
			}
			defer conn.Close()
		})
		ws.GET("/", h.GetStatus)
		ws.GET("/confirm", h.ConfirmDelivery)
	}

	return router
}
