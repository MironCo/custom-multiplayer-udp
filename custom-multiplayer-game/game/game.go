package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
)

// Game implements ebiten.Game interface.
type Game struct {
	gameObjectManager *GameObjectManager
	initialized       bool
	state             GameState
	joinButtonBounds  struct{ x, y, w, h int }
	networkClient     *NetworkClient
}

// Initialize sets up the game objects
func (g *Game) Initialize() {
	if g.initialized {
		return
	}

	g.gameObjectManager = NewGameObjectManager()
	g.networkClient = NewNetworkClient()
	g.state = StateMenu

	// Set up join button bounds (centered on screen)
	g.joinButtonBounds.x = 120
	g.joinButtonBounds.y = 100
	g.joinButtonBounds.w = 80
	g.joinButtonBounds.h = 40

	g.initialized = true
}

func (g *Game) JoinServer() {
	// Connect to server (placeholder URL)
	err := g.networkClient.ConnectToServer()
	if err != nil {
		log.Printf("Failed to connect to server: %v", err)
		return
	}

	err = g.networkClient.JoinGame()
	if err != nil {

	}
	g.state = StatePlaying

}

func (g *Game) AddSoloPlayer() {
	// Create a player at the center of the screen when joining
	player := NewPlayer(160, 120, 2.0, 20, color.RGBA{255, 0, 0, 255}) // Red square
	g.gameObjectManager.AddGameObject(player)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.Initialize()

	switch g.state {
	case StateMenu:
		// Handle join button click
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if x >= g.joinButtonBounds.x && x <= g.joinButtonBounds.x+g.joinButtonBounds.w &&
				y >= g.joinButtonBounds.y && y <= g.joinButtonBounds.y+g.joinButtonBounds.h {
				g.JoinServer()
			}
		}
	case StatePlaying:
		// Update all game objects automatically
		return g.gameObjectManager.UpdateAll()
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	g.Initialize()

	// Clear screen with black background
	screen.Fill(color.RGBA{0, 0, 0, 255})

	switch g.state {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying:
		// Draw all game objects automatically
		g.gameObjectManager.DrawAll(screen)
	}
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	// Draw join button
	buttonColor := color.RGBA{100, 100, 100, 255}
	x, y := ebiten.CursorPosition()
	if x >= g.joinButtonBounds.x && x <= g.joinButtonBounds.x+g.joinButtonBounds.w &&
		y >= g.joinButtonBounds.y && y <= g.joinButtonBounds.y+g.joinButtonBounds.h {
		buttonColor = color.RGBA{150, 150, 150, 255} // Hover color
	}

	// Draw button background
	for i := 0; i < g.joinButtonBounds.h; i++ {
		for j := 0; j < g.joinButtonBounds.w; j++ {
			screen.Set(g.joinButtonBounds.x+j, g.joinButtonBounds.y+i, buttonColor)
		}
	}

	// Draw button text (simple pixel text)
	g.drawSimpleText(screen, "JOIN", g.joinButtonBounds.x+25, g.joinButtonBounds.y+15, color.White)
}

func (g *Game) drawSimpleText(screen *ebiten.Image, text string, x, y int, col color.Color) {
	// Simple bitmap font - just draw "JOIN" manually for now
	if text == "JOIN" {
		// J
		for i := 0; i < 10; i++ {
			screen.Set(x+i, y, col)
			screen.Set(x+8, y+i, col)
			if i > 5 {
				screen.Set(x, y+i, col)
			}
		}
		screen.Set(x, y+9, col)
		screen.Set(x+1, y+9, col)
		screen.Set(x+2, y+9, col)

		// O
		x += 12
		for i := 1; i < 9; i++ {
			screen.Set(x, y+i, col)
			screen.Set(x+6, y+i, col)
		}
		for i := 1; i < 6; i++ {
			screen.Set(x+i, y, col)
			screen.Set(x+i, y+9, col)
		}

		// I
		x += 10
		for i := 0; i < 10; i++ {
			screen.Set(x+2, y+i, col)
		}
		for i := 0; i < 5; i++ {
			screen.Set(x+i, y, col)
			screen.Set(x+i, y+9, col)
		}

		// N
		x += 8
		for i := 0; i < 10; i++ {
			screen.Set(x, y+i, col)
			screen.Set(x+6, y+i, col)
			if i < 8 {
				screen.Set(x+1+i, y+i+1, col)
			}
		}
	}
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
