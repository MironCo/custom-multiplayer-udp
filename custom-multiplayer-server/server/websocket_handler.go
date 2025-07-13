package server

import (
	"custom-multiplayer-server/game"
	"custom-multiplayer-server/types"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	connection  *websocket.Conn
	GameHandler *game.GameHandler
}

func CreateWebsocketHandler(gameHandler *game.GameHandler) *WebsocketHandler {
	return &WebsocketHandler{}
}

func (h *WebsocketHandler) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	var err error
	h.connection, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error: ", err)
	}
	defer h.Close()

	for {
		_, message, err := h.connection.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
		}

		var msg types.WebsocketMessage
		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println("Parse error: ", err)
			continue
		}

		switch msg.MessageType {
		case "join_lobby":
			// do seomthing here
		case "leave_lobby":
			// do soemthing here
		}
	}
}

func (h *WebsocketHandler) Close() {
	h.connection.Close()
}
