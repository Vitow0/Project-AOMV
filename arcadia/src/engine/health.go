package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)
type Monster struct {
    Sprite    rl.Texture2D
    Health    int
    MaxHealth int
}


func (e *Engine) DrawHealthBar() {
	
	barWidth := float32(200)  // defined the height of the healthbar
	barHeight := float32(20)
	barX := float32(20)
	barY := float32(20)

	
	healthPercent := float32(e.Player.Health) / float32(e.Player.MaxHealth) // 

	
	rl.DrawRectangle(int32(barX), int32(barY), int32(barWidth), int32(barHeight), rl.Red) // rectangle of the healthbar red

	
	rl.DrawRectangle(int32(barX), int32(barY), int32(barWidth*healthPercent), int32(barHeight), rl.Green) // rectangle of the healthbar green

	
	healthText := fmt.Sprintf("Health: %d/%d", e.Player.Health, e.Player.MaxHealth)
	rl.DrawText(healthText, int32(barX), int32(barY)-20, 20, rl.RayWhite)

	if e.Player.Health <= 0 && e.StateMenu != HOME {
        e.Player.IsAlive = false
        e.StateEngine = GAMEOVER
    }

    // Test
    if rl.IsKeyPressed(rl.KeyK) {
        damage := 5
        e.Player.Health -= damage
    }
}

// Reset Game if the player die (it is for removing the GameOver screen when you want to go back in the game)
func (e *Engine) ResetGame() {
    if rl.IsKeyPressed(rl.KeySpace) && e.StateMenu == HOME { // if you are on the menu  and you die, it's reset
        e.Player.Health = e.Player.MaxHealth
        e.Player.IsAlive = true
        e.StateEngine = INGAME
    }
}

func (e *Engine) DrawMonsterHealthBar() {
    barWidth := int32(200)    // Total width of the health bar
    barHeight := int32(20)    // Height of the health bar
    
    // Positioning the health bar at the top-right corner of the screen
    barX := int32(1600)  // 20 pixels padding from the right edge
    barY := int32(40)                            // 20 pixels padding from the top

    // Calculate health percentage
    healthPercentage := float32(e.Monster.Health) / float32(e.Monster.MaxHealth)
    currentBarWidth := int32(healthPercentage * float32(barWidth))

    // Draw empty health bar
    rl.DrawRectangle(barX, barY, barWidth, barHeight, rl.Green)

    // Draw current health 
    rl.DrawRectangle(barX, barY, currentBarWidth, barHeight, rl.Red)

    // draw current health text
    healthText := fmt.Sprintf("Health: %d/%d", e.Monster.Health, e.Monster.MaxHealth)
    rl.DrawText(healthText, int32(barX), int32(barY)-20, 20, rl.RayWhite)
}


