package engine

import (
	"main/src/entity"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Init variabel of engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)

	// Init composants of the game
	e.InitEntities()
	e.InitCamera()
	e.InitMusic()
	e.InitMap("textures/map/tilesets/map.json")
	e.InitInventory()

}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 600, Y: 350},
		Health:    100,
		MaxHealth: 100,
		Money:     1000,
		Speed:     2,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite:       e.Player.Sprite,
		CombatSprite: e.Player.CombatSprite,
	}

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "claude",
		Position: rl.Vector2{X: 700, Y: 500},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Slime1",
		Position: rl.Vector2{X: 640, Y: 260},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Slime/Slime-idle.PNG"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "drae",
		Position: rl.Vector2{X: 540, Y: 2400},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/drae/drae.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "drae2",
		Position: rl.Vector2{X: 640, Y: 2690},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/drae/drae.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "drae1",
		Position: rl.Vector2{X: 840, Y: 2500},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/drae/drae.png"),
	})

	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "Slime1",
		Position: rl.Vector2{X: 640, Y: 1560},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/Slime/Slime-idle.PNG"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "scale-a-ton1",
		Position: rl.Vector2{X: 620, Y: 930},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/scale-a-ton/skeleton-idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "scale-a-ton2",
		Position: rl.Vector2{X: 540, Y: 1060},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/scale-a-ton/skeleton-idle.png"),
	})
	e.Player.Money = 12
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "scale-a-ton3",
		Position: rl.Vector2{X: 740, Y: 1160},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/scale-a-ton/skeleton-idle.png"),
	})
	e.Monsters = append(e.Monsters, entity.Monster{
		Name:     "scale-a-ton",
		Position: rl.Vector2{X: 740, Y: 960},
		Health:   45,
		Damage:   9,
		Loot:     []item.Item{}, //items : Sword or shield
		Worth:    17,
		IsAlive:  true,
		Sprite:   rl.LoadTexture("textures/entities/scale-a-ton/skeleton-idle.png"),
	})

	e.Player.Money = 12

}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( // init Camera
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

func (engine *Engine) InitInventory() {
	appleImage := rl.LoadTexture("textures/Menu/apple_regular.png") // Load the apple texture
	if appleImage.ID == 0 {
		rl.DrawText("Failed to load apple texture", int32(20), int32(20), 20, rl.RayWhite)
	}
	engine.AppleTexture = appleImage

	swordImage := rl.LoadTexture("textures/Menu/Sword.png") // Load the sword texture
	if swordImage.ID == 0 {
		rl.DrawText("Failed to load sword texture", int32(20), int32(20), 20, rl.RayWhite)
	}
	engine.SwordTexture = swordImage

	shieldImage := rl.LoadTexture("textures/Menu/Shield.png") // Load the shield texture
	if shieldImage.ID == 0 {
		rl.DrawText("Failed to load shield texture", int32(20), int32(20), 20, rl.RayWhite)
	}
	engine.ShieldTexture = shieldImage

	engine.InitInventory2() // Initialize inventory
}
