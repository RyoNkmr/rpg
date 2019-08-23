package enemy

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/effect"
)

type stats struct {
	name     string
	hp       actor.Hp
	hpmax    actor.Hp
	sp       actor.Sp
	spmax    actor.Sp
	exp      actor.Exp
	effects  effect.EffectMap
	immunity effect.EffectMap
}

type Enemy interface {
	actor.Actor
	GetExp() actor.Exp
}

func (s *stats) GetEffects() []effect.Effect {
	return s.effects.AsOrderedList()
}

func (s *stats) AddEffect(e effect.Effect) {
	s.effects.Add(e)
}

func (s *stats) RemoveEffect(e effect.Effect) {
	s.effects.Remove(e)
}

func (s *stats) BeforeAttack() (messages []actor.Message, isDead bool) {
	// not implemented
	return messages, false
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

func (e *stats) GetExp() actor.Exp {
	return e.exp
}
