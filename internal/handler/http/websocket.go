package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	model "notify-service/internal/handler/model/http"
)

func (h handler) GetStatus(c *gin.Context) {
	ws, ok := c.Get("conn")
	switch ws.(type) {
	case *websocket.Conn:
		if !ok {
			c.JSONP(http.StatusInternalServerError, "Websocket connection not found")
			return
		}

		err := ws.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(model.GetStatus))
		if err != nil {
			c.JSONP(http.StatusInternalServerError, "Error while sending message to client: "+err.Error())
			return
		}

		_, response, err := ws.(*websocket.Conn).ReadMessage()
		if err != nil {
			c.JSONP(http.StatusInternalServerError, "Error while reading message from client: "+err.Error())
			return
		}

		c.JSONP(http.StatusOK, gin.H{"response": string(response)})

	case nil:
		c.JSONP(http.StatusInternalServerError, "Websocket connection is nil")
		return

	default:
		c.JSONP(http.StatusInternalServerError, "Invalid websocket connection")
		return
	}

}

func (h handler) ConfirmDelivery(c *gin.Context) {
	ws, ok := c.Get("conn")
	switch ws.(type) {
	case *websocket.Conn:
		if !ok {
			c.JSONP(http.StatusInternalServerError, "Websocket connection not found")
			return
		}

		err := ws.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(model.GetStatus))
		if err != nil {
			c.JSONP(http.StatusInternalServerError, "Error while sending message to client: "+err.Error())
			return
		}

		c.JSONP(http.StatusOK, gin.H{"response": ""})

	case nil:
		c.JSONP(http.StatusInternalServerError, "Websocket connection is nil")
		return

	default:
		c.JSONP(http.StatusInternalServerError, "Invalid websocket connection")
		return
	}

}
