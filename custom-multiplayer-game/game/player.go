package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Player represents a movable square
type Player struct {
	*BaseGameObject
	speed float64
	size  float64
	color color.Color
}

func NewPlayer(x, y, speed, size float64, playerColor color.Color) *Player {
	return &Player{
		BaseGameObject: NewBaseGameObject(x, y),
		speed:          speed,
		size:           size,
		color:          playerColor,
	}
}

// Update handles player input and movement
func (p *Player) Update() error {
	// Call base update first
	if err := p.BaseGameObject.Update(); err != nil {
		return err
	}

	// Handle movement input
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		p.transform.X -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		p.transform.X += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		p.transform.Y -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		p.transform.Y += p.speed
	}

	// Keep player within screen bounds (assuming 320x240 logical screen)
	if p.transform.X < 0 {
		p.transform.X = 0
	}
	if p.transform.X > 320-p.size {
		p.transform.X = 320 - p.size
	}
	if p.transform.Y < 0 {
		p.transform.Y = 0
	}
	if p.transform.Y > 240-p.size {
		p.transform.Y = 240 - p.size
	}

	return nil
}

// Draw renders the player square
func (p *Player) Draw(screen *ebiten.Image) {
	// Draw a filled rectangle representing the player
	vector.DrawFilledRect(
		screen,
		float32(p.transform.X),
		float32(p.transform.Y),
		float32(p.size),
		float32(p.size),
		p.color,
		false,
	)
}

// Getter methods for accessing player properties
func (p *Player) GetSpeed() float64 {
	return p.speed
}

func (p *Player) SetSpeed(speed float64) {
	p.speed = speed
}

func (p *Player) GetSize() float64 {
	return p.size
}

func (p *Player) GetPosition() (float64, float64) {
	return p.transform.X, p.transform.Y
}

func (p *Player) SetPosition(x, y float64) {
	p.transform.X = x
	p.transform.Y = y
}

// IsKeyJustPressed checks if a key was just pressed this frame
func (p *Player) IsKeyJustPressed(key ebiten.Key) bool {
	return inpututil.IsKeyJustPressed(key)
}
