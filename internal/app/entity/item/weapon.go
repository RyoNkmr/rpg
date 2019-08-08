package item

import "github.com/RyoNkmr/rpg/internal/app/entity"

type weapon struct {
	atkDice *entity.Dice
	baseAtk uint64
}

type Weapon interface {
	getAttack() uint64
}

func (e *equipment) getAttack() (*entity.Dice, uint64) {
	return e.atkDice, e.baseAtk
}
