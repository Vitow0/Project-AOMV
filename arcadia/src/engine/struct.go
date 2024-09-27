package engine

import (
	"main/src/entity"
	"main/src/fight"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	INGAME  engine = iota
	PAUSE    engine = iota
	QUEST    engine = iota
	GAMEOVER engine = iota // 
	INFIGHT engine = iota // defined menus
)

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster
	Monster Monster

	Music       rl.Music
	MusicVolume float32

	Sprites map[string]rl.Texture2D
	AppleTexture rl.Texture2D
	ShieldTexture rl.Texture2D // defined the new textures of the items
	SwordTexture rl.Texture2D
	Camera rl.Camera2D
	

	MapJSON MapJSON

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
	FightManager *fight.FightManager  // defined the fight
	Health     int  // defined health
    MaxHealth  int  // defined the max health of the player
	Inventory     Inventory // defined inventory
    InventoryOpen bool
	Filter        string // defined the filter
	SelectedItem int // defined the selected item
	HPMessage         string  
    HPMessageDuration float32 // defined timer
    HPMessageTimer    float32 
}


