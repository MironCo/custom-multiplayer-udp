package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// GameObject interface that all game objects must implement
type GameObject interface {
	Update() error
	Draw(screen *ebiten.Image)
	IsActive() bool
	SetActive(active bool)
}

// Transform represents position, rotation, and scale
type Transform struct {
	X, Y     float64
	Rotation float64
	ScaleX, ScaleY float64
}

// BaseGameObject provides common functionality for all game objects
type BaseGameObject struct {
	transform Transform
	active    bool
}

func NewBaseGameObject(x, y float64) *BaseGameObject {
	return &BaseGameObject{
		transform: Transform{
			X: x, Y: y,
			Rotation: 0,
			ScaleX: 1, ScaleY: 1,
		},
		active: true,
	}
}

func (b *BaseGameObject) GetTransform() *Transform {
	return &b.transform
}

func (b *BaseGameObject) IsActive() bool {
	return b.active
}

func (b *BaseGameObject) SetActive(active bool) {
	b.active = active
}

// Default implementations that can be overridden
func (b *BaseGameObject) Update() error {
	return nil
}

func (b *BaseGameObject) Draw(screen *ebiten.Image) {
	// Base implementation does nothing
}