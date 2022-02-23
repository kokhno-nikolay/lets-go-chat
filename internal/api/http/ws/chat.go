package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *WSHandler) Chat(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	h.services.ActiveUsers.Add()
	defer ws.Close()
	defer h.services.ActiveUsers.Remove()

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("New message: ", string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err.Error())
			return
		}
	}
}
