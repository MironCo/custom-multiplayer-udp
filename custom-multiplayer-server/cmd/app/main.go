package main

import (
	"custom-multiplayer-server/game"
	"custom-multiplayer-server/server"
)

func main() {
	gameHandler := game.CreateNewGameHandler()

	websocketHandler := server.CreateWebsocketHandler(gameHandler)
	defer websocketHandler.Close()

	serverHandler := server.CreateServerHandler()
	serverHandler.Start()
}
