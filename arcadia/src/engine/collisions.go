package engine

import (
	"main/src/entity"
	"main/src/fight"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (e *Engine) CheckCollisions() {
	for _, monster := range e.Monsters {
		if e.CollisionWithPlayer(monster) {               
			rl.DrawText("FIGHT [E] to engage", int32(rl.GetScreenWidth())/2-rl.MeasureText("FIGHT [E] to engage", 20)/2, int32(rl.GetScreenHeight())/2+50, 20, rl.RayWhite)
			if rl.IsKeyPressed(rl.KeyE) { // if the player hit the button 'E'
				e.StartFight(&e.Player, &monster)
			}
		}
	}
}

func (e *Engine) CollisionWithPlayer(monster entity.Monster) bool { 
	return monster.Position.X > e.Player.Position.X-20 &&
		monster.Position.X < e.Player.Position.X+20 &&
		monster.Position.Y > e.Player.Position.Y-20 &&   // hitbox for starting a fight by pressing letter 'E'
		monster.Position.Y < e.Player.Position.Y+20
}

func (e *Engine) StartFight(player *entity.Player, monster *entity.Monster) {
	e.FightManager = fight.NewFightManager(player, monster) // init the fight
	e.StateEngine = INFIGHT // screen of fight
}
