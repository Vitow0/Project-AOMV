package engine

import (
	"main/src/entity"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) Rendering() {
	rl.DrawTexture(e.Sprites["background"], 0, 0, rl.RayWhite)
}

func (e *Engine) HomeRendering() { // Menu Principale
	rl.DrawTexture(e.Sprites["background"], 0, 0, rl.RayWhite)

	rl.DrawText("ARCADIA", int32(rl.GetScreenWidth())/2-rl.MeasureText("ARCADIA", 155)/2, int32(rl.GetScreenHeight())/2-190, 150, rl.RayWhite)
	rl.DrawText("[Space] to Play", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Space] to Play", 30)/2, int32(rl.GetScreenHeight())/2, 30, rl.RayWhite)
	rl.DrawText("[Escape] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Quit", 290)/2, int32(rl.GetScreenHeight())/2-500, 30, rl.RayWhite)
	rl.DrawText("[K] for Keybinds", int32(rl.GetScreenWidth())/2-rl.MeasureText("[K] for Keybinds", 30)/2, int32(rl.GetScreenHeight())/2+50, 30, rl.RayWhite)
}

func (e *Engine) SettingsRendering() { // Menu Keybinds
	rl.ClearBackground(rl.Blue)

	rl.DrawText("[Escape] to Go Back", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Go Back", 215)/2, int32(rl.GetScreenHeight())/2-500, 30, rl.RayWhite)
	rl.DrawText("Key Binds", int32(rl.GetScreenWidth())/2-rl.MeasureText("Key Binds", 90)/2, int32(rl.GetScreenHeight())/2-400, 90, rl.RayWhite)
	
	rl.DrawText("[U] to Stop music", int32(rl.GetScreenWidth())/2-rl.MeasureText("[U] to Stop music", 81)/2, int32(rl.GetScreenHeight())/2-200, 30, rl.RayWhite)
	rl.DrawText("[I] to Play music", int32(rl.GetScreenWidth())/2-rl.MeasureText("[I] to Play music", 87)/2, int32(rl.GetScreenHeight())/2-150, 30, rl.RayWhite)
	rl.DrawText("[P] Show main quest", int32(rl.GetScreenWidth())/2-rl.MeasureText("[P] Show main quest", 70)/2, int32(rl.GetScreenHeight())/2-100, 30, rl.RayWhite)

	rl.DrawText("[E] to Fight", int32(rl.GetScreenWidth())/2-rl.MeasureText("[E] to Fight", 120)/2, int32(rl.GetScreenHeight())/2, 30, rl.RayWhite)

	rl.DrawText("[W] Go Foraward", int32(rl.GetScreenWidth())/2+110, int32(rl.GetScreenHeight())/2-200, 30, rl.RayWhite)
	rl.DrawText("[S] Go Backward", int32(rl.GetScreenWidth())/2+110, int32(rl.GetScreenHeight())/2-150, 30, rl.RayWhite)
	rl.DrawText("[A] Go Left", int32(rl.GetScreenWidth())/2+110, int32(rl.GetScreenHeight())/2-100, 30, rl.RayWhite)
	rl.DrawText("[D] Go Right", int32(rl.GetScreenWidth())/2+110, int32(rl.GetScreenHeight())/2-50, 30, rl.RayWhite)
}

func (e *Engine) InGameRendering() {
	rl.ClearBackground(rl.Gray)

	rl.BeginMode2D(e.Camera) // begin the render camera

	e.RenderMap()

	e.RenderMonsters()
	e.RenderPlayer()

	rl.EndMode2D() // end of the render camera

	rl.DrawText("[Esc] to Pause", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to Pause", 250)/2, int32(rl.GetScreenHeight())/2-500, 30, rl.Black)

}

func (e *Engine) PauseRendering() { // Menu PAUSE
	rl.DrawTexture(e.Sprites["background2"], 0, 0, rl.RayWhite)

	rl.DrawText("Game Paused", int32(rl.GetScreenWidth())/2-rl.MeasureText("Game Paused", 143)/2, int32(rl.GetScreenHeight())/2-150, 151, rl.White)
	rl.DrawText("[Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to resume", 40)/2, int32(rl.GetScreenHeight())/2+75, 50, rl.White)
	rl.DrawText("[A] to Quit", int32(rl.GetScreenWidth())/2-rl.MeasureText("[A] to Quit", 360)/2, int32(rl.GetScreenHeight())/2-520, 50, rl.White)

}

func (e *Engine) QuestRendering() { // Menu de la QUEST principale
	rl.DrawTexture(e.Sprites["background3"], 0, 0, rl.RayWhite)

	rl.DrawText("[A] to Close", int32(rl.GetScreenWidth())/2-rl.MeasureText("[A] to Close", 310)/2, int32(rl.GetScreenHeight())/2-500, 30, rl.Black)
	rl.DrawText("Vous venez d arriver sur Arcadia, un monde magique conquis par le mal,", int32(rl.GetScreenWidth())/2-rl.MeasureText("[A] to Close", 260)/2, int32(rl.GetScreenHeight())/2-105, 40, rl.Black)
	rl.DrawText("l ancien roi Ghaldur sera présent pour vous guider, battez vous et survivez,", int32(rl.GetScreenWidth())/2-rl.MeasureText("[A] to Close", 280)/2, int32(rl.GetScreenHeight())/2-40, 40, rl.Black)
	rl.DrawText(" pour pouvoir libérer Arcadia !", int32(rl.GetScreenWidth())/2-rl.MeasureText("[A] to Close", 110)/2, int32(rl.GetScreenHeight())/2+25, 40, rl.Black)
}

func (e *Engine) GameOverRendering() { // screen of death
	rl.DrawTexture(e.Sprites["background4"], 0, 0, rl.RayWhite)

	rl.DrawText("GAME OVER", int32(rl.GetScreenWidth())/2-rl.MeasureText("GAME OVER", 135)/2, int32(rl.GetScreenHeight())/2-180, 150, rl.Red)
	rl.DrawText("[Esc] to resume", int32(rl.GetScreenWidth())/2-rl.MeasureText("[Esc] to resume", 40)/2, int32(rl.GetScreenHeight())/2+75, 50, rl.Red)
}

func (e *Engine) RenderPlayer() {

	rl.DrawTexturePro(
		e.Player.Sprite, rl.NewRectangle(0, 0, 64, 64), rl.NewRectangle(e.Player.Position.X, e.Player.Position.Y, 150, 150), rl.Vector2{X: 0, Y: 0}, 0, rl.White)

}

func (e *Engine) RenderMonsters() {
	for _, monster := range e.Monsters {
		if monster.Name == "scale-a-ton" {
			rl.DrawTexturePro(
				monster.Sprite,
				rl.NewRectangle(0, 0, 60, 60), // 10= zoom on the sprit
				rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
				rl.Vector2{X: 0, Y: 0},
				0,
				rl.White,
			)
		} else {
			if monster.Name == "drae" {
				rl.DrawTexturePro(
					monster.Sprite,
					rl.NewRectangle(5, 0, 75, 70), // 10= zoom on the sprit
					rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
					rl.Vector2{X: 0, Y: 0},
					0,
					rl.White,
				)
			} else {
				if monster.Name == "drae2" {
					rl.DrawTexturePro(
						monster.Sprite,
						rl.NewRectangle(5, 0, 75, 70), // 10= zoom on the sprit
						rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
						rl.Vector2{X: 0, Y: 0},
						0,
						rl.White,
					)
				} else { if monster.Name == "drae2" {
					rl.DrawTexturePro(
						monster.Sprite,
						rl.NewRectangle(5, 0, 75, 70), // 10= zoom on the sprit
						rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
						rl.Vector2{X: 0, Y: 0},
						0,
						rl.White,
					)
				} else {
					if monster.Name == "scale-a-ton2" {
						rl.DrawTexturePro(
							monster.Sprite,
							rl.NewRectangle(0, 0, 60, 60), // 10= zoom on the sprit
							rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
							rl.Vector2{X: 0, Y: 0},
							0,
							rl.White,
						)
					} else {
						if monster.Name == "scale-a-ton3" {
							rl.DrawTexturePro(
								monster.Sprite,
								rl.NewRectangle(0, 0, 60, 60), // 10= zoom on the sprit
								rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
								rl.Vector2{X: 0, Y: 0},
								0,
								rl.White,
							)
						} else { if monster.Name == "scale-a-ton1" {
							rl.DrawTexturePro(
								monster.Sprite,
								rl.NewRectangle(0, 0, 60, 60), // 10= zoom on the sprit
								rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
								rl.Vector2{X: 0, Y: 0},
								0,
								rl.White,
							)
						} else {

							if monster.Name == "Slime" {
								rl.DrawTexturePro(
									monster.Sprite,
									rl.NewRectangle(0, 0, 40, 40), // 10= zoom on the sprit
									rl.NewRectangle(monster.Position.X, monster.Position.Y, 50, 50), //change the size of the mob
									rl.Vector2{X: 0, Y: 0},
									0,
									rl.White,
								)
							} else {
								if monster.Name == "drae1" {
									rl.DrawTexturePro(
										monster.Sprite,
										rl.NewRectangle(5, 0, 75, 70), // 10= zoom on the sprit
										rl.NewRectangle(monster.Position.X, monster.Position.Y, 110, 110), //change the size of the mob
										rl.Vector2{X: 0, Y: 0},
										0,
										rl.White,
									)
								} else { if monster.Name == "Slime1" {
									rl.DrawTexturePro(
										monster.Sprite,
										rl.NewRectangle(0, 0, 40, 40), // 10= zoom on the sprit
										rl.NewRectangle(monster.Position.X, monster.Position.Y, 50, 50), //change the size of the mob
										rl.Vector2{X: 0, Y: 0},
										0,
										rl.White,
									)
								} else {
									if monster.Name == "claude" {
										rl.DrawTexturePro(
											monster.Sprite,
											rl.NewRectangle(0, 0, 40, 40), // 10= zoom on the sprit
											rl.NewRectangle(monster.Position.X, monster.Position.Y, 60, 60), //change the size of the mob
											rl.Vector2{X: 0, Y: 0},
											0,
											rl.White,
										)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
}
}
}


func (e *Engine) RenderDialog(m entity.Monster, sentence string) {
	rl.BeginMode2D(e.Camera)

	rl.DrawText(
		sentence,
		int32(m.Position.X),
		int32(m.Position.Y)+50,
		10,
		rl.RayWhite,
	)

	rl.EndMode2D()
}

func (e *Engine) InfightRendering() { // screen of fight
    rl.DrawTexture(e.Sprites["background5"], 0, 0, rl.RayWhite)

	rl.DrawTextureEx(e.Player.CombatSprite, rl.Vector2{X: 390, Y: 590},0, 5, rl.White)

}

func (e *Engine) InfightRenderingTexte() { // text of fight
    rl.DrawText("[A] to Attack", int32(rl.GetScreenWidth())/2-rl.MeasureText("[A] to Attack", 175)/2, int32(rl.GetScreenHeight())/2+390, 50, rl.White)
}

func (e *Engine) RenderInventory() {
    rl.ClearBackground(rl.DarkGray)

    filteredItems := e.Inventory.GetFilteredItems(e.Filter)

    columnCount := 7      // number of column
    itemWidth := 80       // Width of the inventory slot
    itemHeight := 80      // Height of the inventory slot
    paddingX := 100       // Horizontal space between slots
    paddingY := 100       // Vertical space between slots

    totalTableWidth := columnCount*itemWidth + (columnCount-1)*paddingX
    totalTableHeight := 5*itemHeight + 4*paddingY
    screenWidth := rl.GetScreenWidth()
    screenHeight := rl.GetScreenHeight()

    startX := int32((screenWidth - totalTableWidth) / 2)
    startY := int32((screenHeight - totalTableHeight) / 2)
    // reorganize the items when we use filtered items
    totalItems := len(filteredItems)
    rowCount := (totalItems + columnCount - 1) / columnCount
    if rowCount < 5 {
        rowCount = 5 
    }

    for row := 0; row < rowCount; row++ { 
        for column := 0; column < columnCount; column++ {
            // render to move the yellow rectangl for selected item
            xPos := startX + int32(column)*(int32(itemWidth)+int32(paddingX))
            yPos := startY + int32(row)*(int32(itemHeight)+int32(paddingY))
            itemIndex := row*columnCount + column

            if itemIndex == e.SelectedItem {
                rl.DrawRectangleLines(xPos, yPos, int32(itemWidth), int32(itemHeight), rl.Yellow) // draw the yellow rectangle for selected items
            } else {
                rl.DrawRectangleLines(xPos, yPos, int32(itemWidth), int32(itemHeight), rl.RayWhite)
            }

            if itemIndex < totalItems {
                item := filteredItems[itemIndex]

                
                    // Scale the texture to fit in the slot
                    textureScale := float32(itemWidth-20) / float32(item.Image.Width)
                    rl.DrawTextureEx(item.Image, rl.Vector2{X: float32(xPos + 10), Y: float32(yPos + 10)}, 0.0, textureScale, rl.White)
                

                rl.DrawText(fmt.Sprintf("%d. %s", itemIndex+1, item.Name), xPos+80, yPos+20, 20, rl.RayWhite)
            }
        }
    }

    if e.HPMessageTimer < e.HPMessageDuration {  // Display the message with the timer in the used item apple (go see inventory.go)
        e.HPMessageTimer += rl.GetFrameTime()
        rl.DrawText(e.HPMessage, 20, 100, 30, rl.RayWhite)
    }
    e.DrawHealthBar()																// Draw the texte and the health
    rl.DrawText("                                                     Press 'I' to close inventory     'Up/Down/Left/Right' to select     'J' to use item", int32(20), int32(20), 20, rl.RayWhite)
}
func (e *Engine) RenderFightScreen() {
    // Draw the monster sprite
        monsterX := int32(rl.GetScreenWidth()/2) - int32(e.Monster.Sprite.Width)/2
        monsterY := int32(rl.GetScreenHeight()/2) - int32(e.Monster.Sprite.Height)/2
        rl.DrawTexture(e.Monster.Sprite, monsterX, monsterY, rl.White)
    
    // Draw the monster's health bar
    e.DrawMonsterHealthBar()
	// Draw the player healthbar
	e.DrawHealthBar()
}