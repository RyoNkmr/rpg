package enemy

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
)

type stats struct {
	name  string
	hp    actor.Hp
	hpmax actor.Hp
	sp    actor.Sp
	spmax actor.Sp
}

type Enemy interface {
	actor.Actor
}

func (s *stats) IsFriend() bool {
	return false
}

func (s *stats) IsAlive() bool {
	return s.hp > 0
}

func (e *stats) Damage(d actor.Damage) (message string, isDead bool) {
	message = fmt.Sprintf("%s takes %d damage", e.GetName(), d)
	e.hp -= int64(d)
	return message, !e.IsAlive()
}

func (e *stats) GetName() string {
	return e.name
}

func (e *stats) GetStatsString() string {
	return fmt.Sprintf("%s: Hp: %d", e.name, e.hp)
}

func (e *stats) GetStats() (hp, maxHp actor.Hp, sp, maxSp actor.Sp, hunger actor.Hunger) {
	return e.hp, e.hpmax, e.sp, e.spmax, 0
}
