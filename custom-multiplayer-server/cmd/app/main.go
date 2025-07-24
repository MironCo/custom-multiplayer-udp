package main

import (
	"custom-multiplayer-server/game"
	"custom-multiplayer-server/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	gameHandler := game.CreateNewGameHandler()

	websocketHandler := server.CreateWebsocketHandler(gameHandler)

	serverHandler := server.CreateServerHandler(websocketHandler)

	http.HandleFunc("/ws", websocketHandler.HandleWebsocket)

	go func() {
		fmt.Println("HTTP Server listing on Port http://127.0.0.1:8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatal("HTTP server failed:", err)
		}
	}()

	serverHandler.Start()
}
