package race

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
)

type race struct {
	strength     uint64
	intelligence uint64
	dexterity    uint64
	constitution uint64
	hitDiceBase  dice.DiceSide

	level   actor.Level
	exp     actor.Exp
	expRate float64

	levelTable []actor.Exp
}

type Race interface {
	fmt.Stringer
	GetExp() actor.Exp
	GainExp(actor.Exp) (isLevelChanged bool)
	LoseExp(actor.Exp) (isLevelChanged bool)
	GetCurrentLevel() actor.Level
	GetExpToNextLevel() actor.Exp
	GetHitDiceBase() dice.DiceSide
}

func newRace(str, intl, dex, con uint64, expRate float64, hitDiceBase dice.DiceSide) *race {
	r := &race{
		strength:     str,
		intelligence: intl,
		dexterity:    dex,
		constitution: con,
		expRate:      expRate,
		hitDiceBase:  hitDiceBase,
		level:        1,
	}
	return r
}

func (r *race) GetExp() actor.Exp {
	return r.exp
}

func (r *race) GainExp(exp actor.Exp) (isLevelChanged bool) {
	cl := r.GetCurrentLevel()
	r.updateLevel(exp)
	return cl != r.GetCurrentLevel()
}

func (r *race) LoseExp(exp actor.Exp) (isLevelChanged bool) {
	cl := r.GetCurrentLevel()
	r.updateLevel(-exp)
	return cl != r.GetCurrentLevel()
}

func (r *race) updateLevel(exp actor.Exp) {
	r.exp += exp

	if exp > 0 && r.GetExpToNextLevel() <= 0 {
		for i := r.level - 1; ; i++ {
			if r.exp >= expTable[i] {
				r.level = i + 2
				return
			}
		}
	}

	if exp < 0 && expTable[r.level-1] > r.exp {
		for i := r.level - 1; ; i-- {
			if r.exp >= expTable[i] {
				r.level = i + 2
				return
			}
		}
	}
}

func (r *race) GetCurrentLevel() (l actor.Level) {
	return r.level
}

func (r *race) GetExpToNextLevel() actor.Exp {
	return expTable[r.level-1] - r.exp
}

func (r *race) GetHitDiceBase() dice.DiceSide {
	return r.hitDiceBase
}
