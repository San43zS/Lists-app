package websocket

import (
	"Lists-app/internal/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Websocket struct {
	Ws *websocket.Conn
}

type Clients struct {
	srv service.Service
}

func New(service service.Service) *Clients {

	return &Clients{
		srv: service,
	}
}

var upgraded = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)

func broadcastMessage(message string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Ошибка при отправке сообщения:", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func (ws *Websocket) handleWebSocket(w http.ResponseWriter, r *http.Request) http.Handler {
	conn, err := upgraded.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка при подключении WebSocket:", err)
		return nil
	}
	defer conn.Close()

	clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}
	}
	return nil
}

func Start() {
	//http.HandleFunc("/ws", handleWebSocket)

}
