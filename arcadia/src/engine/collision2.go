package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const tileSize = 32 //size of  one tile

func (e *Engine) CheckTileCollisionsAt(newPos rl.Vector2) bool {

	playerTileX := int(newPos.X) / tileSize // hitbox of the player
	playerTileY := int(newPos.Y) / tileSize

	if playerTileY >= 0 && playerTileY < len(tileMap) && playerTileX >= 0 && playerTileX < len(tileMap[playerTileY]) { // if tile is zero
		if tileMap[playerTileY][playerTileX] == 1 { // if the tile is one
			return true // collision
		}
	}
	return false // tile traversable
}
