package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func configuration(c *gin.Context, upgrader *websocket.Upgrader) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return nil, err
	}
	c.Set("conn", ws)

	return ws, nil
}
