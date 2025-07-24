package server

import (
	"custom-multiplayer-server/game"
	"custom-multiplayer-server/types"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	GameHandler *game.GameHandler
}

func CreateWebsocketHandler(gameHandler *game.GameHandler) *WebsocketHandler {
	return &WebsocketHandler{
		GameHandler: gameHandler,
	}
}

func (h *WebsocketHandler) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	var err error
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error: ", err)
	}
	defer connection.Close()

	for {
		_, message, err := connection.ReadMessage()
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
			var joinMessage types.JoinLobbyMessage

			if err := json.Unmarshal(msg.MessageData, &joinMessage); err != nil {
				log.Println("Error parsing join_lobby data:", err)
				continue
			}

			udpAddr, err := net.ResolveUDPAddr("udp", joinMessage.UDPAddress)
			if err != nil {
				fmt.Println("Issue resolving UDP address")
				continue
			}

			h.GameHandler.AddPlayerToRandomGame(udpAddr)
		case "leave_lobby":
			// do soemthing here
		}
	}
}
