package engine

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {

    rl.SetTargetFPS(60)

    for engine.IsRunning {

        rl.BeginDrawing()

        // Call function inventory
        engine.HandleInventory()

        if !engine.InventoryOpen {
            // Gestion of the game
            switch engine.StateMenu {
            case HOME:
                engine.HomeRendering()
                engine.HomeLogic()

            case SETTINGS:
                engine.SettingsRendering()
                engine.SettingsLogic()

            case PLAY:
                switch engine.StateEngine {
                case INGAME:
                    engine.InGameRendering()
                    engine.InGameLogic()

                case PAUSE:
                    engine.PauseRendering()
                    engine.PauseLogic()

                case QUEST:
                    engine.QuestRendering()
                    engine.QuestLogic()

                case GAMEOVER:
                    engine.GameOverRendering()
                    engine.GameOverLogic()

                case INFIGHT:
                    engine.InfightRendering()
                    engine.InfightRenderingTexte()
                    engine.InFightLogic()
                    engine.RenderFightScreen()
                }
            }
        }

        rl.EndDrawing()
    }
}