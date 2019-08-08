package item

import "github.com/RyoNkmr/rpg/internal/app/entity"

type equipment struct {
	atkDice *entity.Dice
	baseAtk uint64
	baseInt uint64
	baseDex uint64
	baseCon uint64
}

func (e *equipment) getAttackEffect() uint64 {
	return e.atkDice.Cast()
}
