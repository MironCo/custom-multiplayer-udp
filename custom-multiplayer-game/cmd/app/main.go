package main

import "custom-multiplayer-game/game"

func main() {
	game := game.StartGame()
	game.Update()
}
