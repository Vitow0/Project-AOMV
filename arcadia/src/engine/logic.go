package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) HomeLogic() {

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Pixel 2.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)
	if rl.IsKeyPressed(rl.KeyU) { // Stop Music
		rl.SetMusicVolume(e.Music, 0)
	}

	if rl.IsKeyPressed(rl.KeyY) { // Play Music
		rl.SetMusicVolume(e.Music, 1)
	}

	//Menus
	if rl.IsKeyPressed(rl.KeySpace) {
		e.StateMenu = PLAY
		e.StateEngine = INGAME
		rl.StopMusicStream(e.Music)

	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.IsRunning = false
	}
	if rl.IsKeyPressed(rl.KeyK) {
		e.StateMenu = SETTINGS
	}
}

func (e *Engine) SettingsLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}
	//Musique
	rl.UpdateMusicStream(e.Music)
}
func (e *Engine) GameOverLogic() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateMenu = HOME
	}
	if rl.IsKeyPressed(rl.KeyU) { // Stop Music
		rl.SetMusicVolume(e.Music, 0)
	}

	if rl.IsKeyPressed(rl.KeyY) { // Play Music
		rl.SetMusicVolume(e.Music, 1)
	}
}

func (e *Engine) InGameLogic() {
	
	

	// Camera
	e.Camera.Target = rl.Vector2{X: e.Player.Position.X + 70, Y: e.Player.Position.Y + 70}
	e.Camera.Offset = rl.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = PAUSE
	}
	if rl.IsKeyPressed(rl.KeyP) { // Display the quest 
		e.StateEngine = QUEST
	}

	if rl.IsKeyPressed(rl.KeyN) {
		e.StateEngine = GAMEOVER
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		e.StateEngine = INFIGHT
	}
		//health
		if e.Player.Health > 100 {
			e.Player.Health = 100
		}
	
	
	//collisions
	e.CheckCollisions()
	//health
	e.DrawHealthBar()
	//reset le jeu
	e.ResetGame()
    
	

	//Musique
	if !rl.IsMusicStreamPlaying(e.Music) {
		e.Music = rl.LoadMusicStream("sounds/music/Pixel 1.mp3")
		rl.PlayMusicStream(e.Music)
	}
	rl.UpdateMusicStream(e.Music)

	if rl.IsKeyPressed(rl.KeyU) { // Pause Music
		rl.SetMusicVolume(e.Music, 0)
	}

	if rl.IsKeyPressed(rl.KeyY) { // Play Music
		rl.SetMusicVolume(e.Music, 1)
	}

	//collision of the map

	newPosition := e.Player.Position

    
    if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
        newPosition.Y -= e.Player.Speed
        if !e.CheckTileCollisionsAt(newPosition) {
            e.Player.Position.Y = newPosition.Y
        }
    }
    if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
        newPosition.Y += e.Player.Speed
        if !e.CheckTileCollisionsAt(newPosition) {
            e.Player.Position.Y = newPosition.Y
        }
    }

    
    if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
        newPosition.X -= e.Player.Speed
        if !e.CheckTileCollisionsAt(newPosition) {
            e.Player.Position.X = newPosition.X
        }
    }
    if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
        newPosition.X += e.Player.Speed
        if !e.CheckTileCollisionsAt(newPosition) {
            e.Player.Position.X = newPosition.X
        }
    }
}

func (e *Engine) MonsterCollisions() {
	for _, monster := range e.Monsters {
		if monster.Position.X > e.Player.Position.X-30 &&
			monster.Position.X < e.Player.Position.X+30 &&
			monster.Position.Y > e.Player.Position.Y-30 &&
			monster.Position.Y < e.Player.Position.Y+30 {

			// Dialogue pour "claude"
			if monster.Name == "claude" {
				e.NormalTalk(monster, "L'esprit de la forêt d'Arcadia réside en moi...")
			}

			// Dialogue pour "Slime"
			if monster.Name == "Slime" {
				e.NormalTalk(monster, "Oh...\n,Cela fais fo-fo-FORRT longtemps\n que je n'ai pas vue d'humain...\nSurvie a l'épreuve de la foret\n, aide moi a sauver Arcadia...")
				if rl.IsKeyPressed(rl.KeyE) {

				}
			} else if monster.Name == "Slime1" {
				e.NormalTalk(monster, "Intéréssant...\ntu te débrouilles bien HUMAIN...\n Désormais tu dois vaincre les monstres qui occupent mon chateau\n Bonne chance, AhaAHAHah...")
			}
		}
	}
}
func (e *Engine) NormalTalk(m entity.Monster, sentence string) {
	e.RenderDialog(m, sentence)
}

func (e *Engine) PauseLogic() {
	//Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		e.StateEngine = INGAME
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		e.StateMenu = HOME
		rl.StopMusicStream(e.Music)
	}

	//Musique
	rl.UpdateMusicStream(e.Music)
}

func (e *Engine) GAMEOVERLogic() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		e.StateEngine = PAUSE
	}
}

func (e *Engine) QuestLogic() {
	if rl.IsKeyPressed(rl.KeyQ) { // go back in game
		e.StateMenu = PLAY
		e.StateEngine = INGAME
	}
}

func (e *Engine) InFightLogic() {
	// check the fight
	if e.FightManager != nil && e.FightManager.InFight {
		e.FightManager.UpdateFight() // update fight
	} else {
		// fight over, go back in the map
		e.StateEngine = INGAME
	}
}

func (e *Engine) ManageInventory() { // add items to each type and add the filtered 
	
	inventory := Inventory{}

	inventory.AddItem(Item{Name: "Sword", Type: "Weapon", Value: 150})
	inventory.AddItem(Item{Name: "Health", Type: "Apple", Value: 50})
	inventory.AddItem(Item{Name: "Shield", Type: "Armor", Value: 100})

	inventory.DisplayInventory()

	inventory.SortInventory("name")
	inventory.DisplayInventory()

	inventory.SortInventory("type")
	inventory.DisplayInventory()

	inventory.SortInventory("value")
	inventory.DisplayInventory()	
}

func (e *Engine) FullScreen() {
	// put in full screen
    if rl.IsKeyPressed(rl.KeyTab) {
        rl.ToggleFullscreen()
    }
}