package playerclass

import "github.com/RyoNkmr/rpg/internal/app/entity/dice"

type warrior struct {
	hitDiceBonus dice.DiceSide
}

func NewWarrior() *warrior {
	return &warrior{
		hitDiceBonus: 9,
	}
}

func (p *warrior) String() string {
	return "warriror"
}

func (p *warrior) GetHitDiceBonus() dice.DiceSide {
	return p.hitDiceBonus
}
