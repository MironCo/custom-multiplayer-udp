package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game interface.
type Game struct {
	gameObjectManager *GameObjectManager
	initialized       bool
}

// Initialize sets up the game objects
func (g *Game) Initialize() {
	if g.initialized {
		return
	}

	g.gameObjectManager = NewGameObjectManager()
	
	// Create a player at the center of the screen
	player := NewPlayer(160, 120, 2.0, 20, color.RGBA{255, 0, 0, 255}) // Red square
	g.gameObjectManager.AddGameObject(player)
	
	g.initialized = true
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.Initialize()
	
	// Update all game objects automatically
	return g.gameObjectManager.UpdateAll()
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Initialize()
	
	// Clear screen with black background
	screen.Fill(color.RGBA{0, 0, 0, 255})
	
	// Draw all game objects automatically
	g.gameObjectManager.DrawAll(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func StartGame() *Game {
	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Moving Squares")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
	return game
}
