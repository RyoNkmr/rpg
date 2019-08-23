package enemy

import (
	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/effect"
	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
)

type snake struct {
	*stats
	attackDice dice.Dice
}

func NewSnake() *snake {
	d := dice.NewDice(3, 6)
	hp := int64(d.Cast())
	e := &stats{
		hp:      hp,
		name:    "snake",
		exp:     actor.Exp(hp),
		effects: effect.EffectMap{},
	}
	return &snake{
		stats:      e,
		attackDice: dice.NewDice(2, 2),
	}
}

func (e *snake) Attack(t actor.Actor) (d actor.Damage, ms []actor.Message) {
	ms = make([]actor.Message, 0, 2)
	ms = append(ms, "snake bites you.")
	d = 1

	if e.attackDice.Cast() >= 1 {
		ms = append(ms, "snake poisoned you.")
		t.AddEffect(effect.Poisoned)
		d += 1
	}

	if e.attackDice.Cast() == 4 {
		ms = append(ms, "snake crushes you.")
		d += 1
	}

	return d, ms
}
