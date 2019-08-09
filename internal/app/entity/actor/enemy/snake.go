package enemy

import (
	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
)

type snake struct {
	*stats
	attackDice dice.Dice
}

func NewSnake() *snake {
	d := dice.NewDice(3, 6)
	e := &stats{
		Hp:   int64(d.Cast()),
		name: "snake",
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

	if e.attackDice.Cast() >= 3 {
		ms = append(ms, "snake crushes you.")
		d += 2
	}

	return d, ms
}
