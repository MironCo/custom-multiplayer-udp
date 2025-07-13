package main

import (
	"custom-multiplayer-server/game"
	"custom-multiplayer-server/server"
	"log"
	"net/http"
)

func main() {
	gameHandler := game.CreateNewGameHandler()

	websocketHandler := server.CreateWebsocketHandler(gameHandler)

	serverHandler := server.CreateServerHandler(websocketHandler)

	http.HandleFunc("/ws", websocketHandler.HandleWebsocket)

	go func() {
		log.Println("HTTP server starting on :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatal("HTTP server failed:", err)
		}
	}()

	serverHandler.Start()
}
