package entity

import (
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Name     string
	Position rl.Vector2
	Health   int
	MaxHealth int
	Damage   int
	Loot     []item.Item
	Worth    int 
	Defense int
	IsAlive bool

	Sprite rl.Texture2D
	CombatSprite rl.Texture2D
}

func (m *Monster) Attack(player *Player) { // damage of the attack of player and the defnese
	damage := 5 - player.Defense 
	if damage < 0 {
		damage = 0 
	}
	player.Health -= damage
}
func (m *Monster) Defend() { // defense of the monster
	m.Defense += 5 
}

