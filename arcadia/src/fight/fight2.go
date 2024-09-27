package fight

import (
	"main/src/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FightManager struct {
	Player    *entity.Player
	Monster   *entity.Monster
	Turn      fightState
	InFight   bool
	GameOver  bool
}

type fightState int

const (
	PLAYER_TURN fightState = iota
	MONSTER_TURN
)

func NewFightManager(player *entity.Player, monster *entity.Monster) *FightManager {
	return &FightManager{
		Player:  player,
		Monster: monster,
		Turn:    PLAYER_TURN,
		InFight: true,
	}
}

func (fm *FightManager) UpdateFight() {
	// if fight over
	if fm.GameOver || fm.Player.Health <= 0 || fm.Monster.Health <= 0 {
		fm.EndFight()
		return
	}

	switch fm.Turn {
	case PLAYER_TURN:
		// turn of player
		if rl.IsKeyPressed(rl.KeyA) {
			fm.Player.Attack(fm.Monster)
			fm.Turn = MONSTER_TURN
		} else if rl.IsKeyPressed(rl.KeyC) {
			fm.Player.Defend()
			fm.Turn = MONSTER_TURN
		}
	case MONSTER_TURN:
		// turn of monster
		fm.Monster.Attack(fm.Player)
		fm.Turn = PLAYER_TURN
	}
}

func (fm *FightManager) EndFight() { // end of the fight
	if fm.Player.Health <= 0 {
		fm.Player.IsAlive = false
		fm.GameOver = true
	} else if fm.Monster.Health <= 0 {
		fm.Player.Inventory = append(fm.Player.Inventory, fm.Monster.Loot...)
		fm.Player.Money += fm.Monster.Worth
	}
	fm.InFight = false
}
