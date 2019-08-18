package enemy

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
)

type stats struct {
	name string
	Hp   int64
}

type Enemy interface {
	actor.Actor
	// Attack(*player.Player) (damage actor.Damage, message []string)
	// Damage(actor.Damage) (message string, isDead bool)
}

func (e *stats) Damage(d actor.Damage) (message string, isDead bool) {
	message = fmt.Sprintf("%s takes %d damage", e.GetName(), d)
	e.Hp -= int64(d)
	return message, e.Hp <= 0
}

func (e *stats) GetName() string {
	return e.name
}

func (e *stats) GetStats() string {
	return fmt.Sprintf("%s: Hp: %d", e.name, e.Hp)
}
