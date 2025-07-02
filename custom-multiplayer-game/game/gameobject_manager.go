package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// GameObjectManager handles all game objects in the scene
type GameObjectManager struct {
	gameObjects []GameObject
}

func NewGameObjectManager() *GameObjectManager {
	return &GameObjectManager{
		gameObjects: make([]GameObject, 0),
	}
}

// AddGameObject adds a new game object to the manager
func (gom *GameObjectManager) AddGameObject(obj GameObject) {
	gom.gameObjects = append(gom.gameObjects, obj)
}

// RemoveGameObject removes a game object from the manager
func (gom *GameObjectManager) RemoveGameObject(obj GameObject) {
	for i, gameObj := range gom.gameObjects {
		if gameObj == obj {
			// Remove by replacing with last element and shortening slice
			gom.gameObjects[i] = gom.gameObjects[len(gom.gameObjects)-1]
			gom.gameObjects = gom.gameObjects[:len(gom.gameObjects)-1]
			break
		}
	}
}

// UpdateAll calls Update on all active game objects
func (gom *GameObjectManager) UpdateAll() error {
	for _, obj := range gom.gameObjects {
		if obj.IsActive() {
			if err := obj.Update(); err != nil {
				return err
			}
		}
	}
	return nil
}

// DrawAll calls Draw on all active game objects
func (gom *GameObjectManager) DrawAll(screen *ebiten.Image) {
	for _, obj := range gom.gameObjects {
		if obj.IsActive() {
			obj.Draw(screen)
		}
	}
}

// GetGameObjectCount returns the number of game objects
func (gom *GameObjectManager) GetGameObjectCount() int {
	return len(gom.gameObjects)
}

// Clear removes all game objects
func (gom *GameObjectManager) Clear() {
	gom.gameObjects = gom.gameObjects[:0]
}