package engine

import (
	"strconv"
	"sort"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Item characteristics in the inventory
type Item struct {
	Name  string
	Type  string  //  "Weapon(attack)", "Health", "Armor(defense)"
	Value int     //  "gold value "
	Image rl.Texture2D
}

// Inventory : a list of items
type Inventory struct {
	Items []Item
}

// AddItem adds a new item
func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
	// Display a message (unofficial)
	message := "Added " + item.Name + " to inventory"
	rl.DrawText(message, 10, 10, 20, rl.RayWhite)
}

// shows all items in the inventory
func (inv *Inventory) DisplayInventory() {
	yPosition := 70
	for _, item := range inv.Items {
		itemInfo := "- " + item.Name + " (Type: " + item.Type + ", Value: " + strconv.Itoa(item.Value) + ")"
		rl.DrawText(itemInfo, 10, int32(yPosition), 20, rl.RayWhite)
		yPosition += 30
	}
}

// sorts items by a chosen criterion
func (inv *Inventory) SortInventory(criteria string) {
	switch criteria {
	case "name":
		sort.SliceStable(inv.Items, func(i, j int) bool {
			return inv.Items[i].Name < inv.Items[j].Name
		})
	case "type":
		sort.SliceStable(inv.Items, func(i, j int) bool {
			return inv.Items[i].Type < inv.Items[j].Type
		})
	case "value":
		sort.SliceStable(inv.Items, func(i, j int) bool {
			return inv.Items[i].Value < inv.Items[j].Value
		})
	default:
		rl.DrawText("Unknown sorting criteria", 10, 10, 20, rl.Red)
		return
	}
	// Display inventory  that got filtred
	rl.DrawText("Inventory sorted by "+criteria, 10, 10, 20, rl.RayWhite)
}

// list of items based on the criteria
func (inv *Inventory) GetFilteredItems(criteria string) []Item {
	// Clone the inventory items
	filteredItems := make([]Item, len(inv.Items))
	copy(filteredItems, inv.Items)

	switch criteria {
	case "name":
		sort.SliceStable(filteredItems, func(i, j int) bool {
			return filteredItems[i].Name < filteredItems[j].Name
		})
	case "type":
		sort.SliceStable(filteredItems, func(i, j int) bool {
			return filteredItems[i].Type < filteredItems[j].Type
		})
	case "value":
		sort.SliceStable(filteredItems, func(i, j int) bool {
			return filteredItems[i].Value < filteredItems[j].Value
		})
	}
	return filteredItems
}
