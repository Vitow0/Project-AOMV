package engine

import (
rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Load() {
// Load character textures
e.Player.Sprite = rl.LoadTexture("textures/entities/Character/Character_Idle.png")
e.Player.CombatSprite = rl.LoadTexture("textures/entities/Character/Character_Walk.png")
//Load monsters texures

// Load menu texture
e.Sprites["background"] = rl.LoadTexture("textures/Menu/Main_menu_background.png")
e.Sprites["background2"] = rl.LoadTexture("textures/Menu/Pause_menu_background.png")
e.Sprites["background3"] = rl.LoadTexture("textures/Menu/Parchemin.png")
e.Sprites["background4"] = rl.LoadTexture("textures/Menu/GameOver.png")
e.Sprites["background5"] = rl.LoadTexture("textures/Menu/Fight_menu_background.png")
// load items of the inventory
e.AppleTexture = rl.LoadTexture("textures/Menu/apple_regular.png")
e.ShieldTexture = rl.LoadTexture("textures/Menu/Shield.png")
e.SwordTexture = rl.LoadTexture("textures/Menu/Sword.png")
}

func (e *Engine) Unload() {
// unload the textures
rl.UnloadTexture(e.Player.Sprite)

for _, sprite := range e.Sprites {
rl.UnloadTexture(sprite)
}

for _, monster := range e.Monsters {
rl.UnloadTexture(monster.Sprite)
}
}
