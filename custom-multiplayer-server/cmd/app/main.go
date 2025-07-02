package main

import (
	"custom-multiplayer-server/server"
)

func main() {
	serverHandler := server.CreateServerHandler()
	serverHandler.Start()
}
