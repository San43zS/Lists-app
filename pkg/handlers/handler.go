package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Endpoints fo registration & authorization
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	// api for working with todo lists & their items
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			// for anyone who has connection with list on id
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")
		}

		items := lists.Group(":id/items")
		{
			items.POST("/")
			items.GET("/")
			items.GET("/:item_id")
			items.PUT("/:item_id")
			items.DELETE("/:item_id")
		}
	}
	return router
}
