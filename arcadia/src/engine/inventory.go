package engine

import (rl "github.com/gen2brain/raylib-go/raylib"
			"math/rand"  
    		"time"   		

)

func (engine *Engine) HandleInventoryInput() {
    if engine.Player.Health <= 0 {
        engine.StateEngine = GAMEOVER // if health = 0, then display Gameover
        return
    }
    columnCount := 7 // show the column in the inventory

    if rl.IsKeyPressed(rl.KeyUp) {                  // key for moving in the inventory
        engine.SelectedItem -= columnCount
        if engine.SelectedItem < 0 {
            engine.SelectedItem = 0
        }
    }
    if rl.IsKeyPressed(rl.KeyDown) {
        engine.SelectedItem += columnCount
        if engine.SelectedItem>= len(engine.Inventory.Items) {
            engine.SelectedItem = len(engine.Inventory.Items) - 1
        }
    }
    if rl.IsKeyPressed(rl.KeyLeft) {
        engine.SelectedItem--
        if engine.SelectedItem < 0 {
            engine.SelectedItem = 0
        }
    }
    if rl.IsKeyPressed(rl.KeyRight) {
        engine.SelectedItem++
        if engine.SelectedItem >= len(engine.Inventory.Items) {
            engine.SelectedItem = len(engine.Inventory.Items) - 1
        }
    }

    // I to close inventory
    if rl.IsKeyPressed(rl.KeyI) {
        engine.Filter = ""  
        engine.RenderInventory()
    }

    // J for using the items
    if rl.IsKeyPressed(rl.KeyJ) {
        engine.UseSelectedItem()
    }

    // key for filters
    if rl.IsKeyPressed(rl.KeyN) {
        engine.Filter = "name"
    }
    if rl.IsKeyPressed(rl.KeyT) {
        engine.Filter = "type"
    }
    if rl.IsKeyPressed(rl.KeyV) {
        engine.Filter = "value"
    }
}
func (e *Engine) HandleInventory() {
    // open or close the inventory
    if rl.IsKeyPressed(rl.KeyI) && e.StateEngine != GAMEOVER {
        e.InventoryOpen = !e.InventoryOpen
    }

    // gestion of the inventory
    if e.InventoryOpen && e.StateEngine != GAMEOVER {
        e.RenderInventory()
        e.HandleInventoryInput()

        // check for health
        if e.Player.Health <= 0 {
            e.StateEngine = GAMEOVER
            e.InventoryOpen = false // close the inventory automatically
        }
    }
}


func (e *Engine) UseSelectedItem() {
    filteredItems := e.Inventory.GetFilteredItems(e.Filter) // go to the function filtered items

    if e.SelectedItem >= 0 && e.SelectedItem < len(filteredItems) { // if item is filtered all items will be moved from another sections
        selectedItem := filteredItems[e.SelectedItem]

        switch selectedItem.Type {
        case "Apple":
            // Just a random number generator (maybe modify rand.seed beacause i updated for 1.20.0 and this variable is becoming useless now)
            rand.Seed(time.Now().UnixNano())

            // Generate random luck : 1/4 chance to lose HP, 3/4 to gain HP
            luck := rand.Intn(4)

            if luck == 0 {
                e.Player.Health -= 10
                e.HPMessage = "Unlucky! Lost 10 HP."
            } else {
                e.Player.Health += 10
                if e.Player.Health > e.Player.MaxHealth {
                    e.Player.Health = e.Player.MaxHealth
                }
                e.HPMessage = "Lucky! Gained 10 HP."
            }

            // Start message timer : 3 seconds (need to import time because the text is only display for almost 0 seconds)
            e.HPMessageDuration = 3.0  
            e.HPMessageTimer = 0.0

            // Remove the used apple from inventory
            e.Inventory.Items = append(e.Inventory.Items[:e.SelectedItem], e.Inventory.Items[e.SelectedItem+1:]...)

            // Update selected item
            if e.SelectedItem >= len(e.Inventory.Items) {
                e.SelectedItem = len(e.Inventory.Items) - 1
            }
        }
    }
}

func (e *Engine) InitInventory2() {
    // Add apples to the inventory
    apple := Item{
        Name:  "Apple",
        Type:  "Apple",
        Value: 20,
        Image: e.AppleTexture,
    }
    for i := 0; i < 4; i++ {
        e.Inventory.AddItem(apple)
    }

    // Add the sword to the inventory
    sword := Item{
        Name:  "Sword",
        Type:  "Weapon",
        Value: 60,
        Image: e.SwordTexture,
    }
    e.Inventory.AddItem(sword)

    // Add the shield to the inventory
    shield := Item{
        Name:  "Shield",
        Type:  "Armor",
        Value: 70,
        Image: e.ShieldTexture,
    }
    e.Inventory.AddItem(shield) // function to add items
}


