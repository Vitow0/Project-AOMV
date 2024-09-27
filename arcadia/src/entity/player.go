package entity

import (
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {

	Position  rl.Vector2
	Health    int
	MaxHealth int
	Money     int
	Speed     float32
	Inventory []item.Item
	Defense int // defined defense

	IsAlive bool

	Sprite rl.Texture2D
	CombatSprite rl.Texture2D
}

func (p *Player) Attack(monster *Monster) { // damage of the player
	damage := 10 - monster.Defense
	if damage < 0 {
		damage = 0 
	}
	monster.Health -= damage
}


func (p *Player) Defend() { // defense of the player
	p.Defense += 5 
}



