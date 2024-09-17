package engine

import (
	"main/src/entity"
	"main/src/item"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 800
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 300, Y: 300},
		Health:    100,
		Money:     1000,
		Speed:     5,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 400, Y: 320},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Slime",
		Position: rl.Vector2{X: 335, Y: 330},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items ==} épée ou armure
		Worth:    17,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Slime/Slime-idle.PNG"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Slime",
		Position: rl.Vector2{X: 385, Y: 330},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items ==} épée ou armure
		Worth:    17,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Slime/Slime-idle.PNG"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Slime",
		Position: rl.Vector2{X: 360, Y: 320},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items ==} épée ou armure
		Worth:    17,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/Slime/Slime-idle.PNG"),
	})

	e.Player.Money = 12
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}

func (e *Engine) InitMusic() {
	rl.InitAudioDevice()

	e.Music = rl.LoadMusicStream("sounds/music/OSC-Ambient-Time-08-Egress.mp3")

	rl.PlayMusicStream(e.Music)
}
func (e *Engine) NewHealtBar () {
	rl.InitWindow(800,600, "ShowHealth")
	defer rl.CloseWindow()
	
	player := entity.Player{Health: 100}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		displayHealth(player.Health)
		rl.EndDrawing()
	}
}

func displayHealth(health int) {
	healthText := "Vie: " + strconv.Itoa(health)
	rl.DrawText(healthText, 10, 10, 20, rl.Red)
}
