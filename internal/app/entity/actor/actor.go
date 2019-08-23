package actor

import "github.com/RyoNkmr/rpg/internal/app/entity/actor/effect"

type Message = string
type Damage = uint64

type Hp = int64
type Sp = int64
type Hunger = int64

type Level = uint16
type Exp = int64

type Actor interface {
	BeforeAttack() (messages []Message, isDead bool)
	Attack(Actor) (Damage, []Message)
	AddEffect(effect.Effect)
	RemoveEffect(effect.Effect)
	GetEffects() []effect.Effect
	Damage(Damage) (message Message, isDead bool)
	IsFriend() bool
	IsAlive() bool
	GetName() string
	GetStats() (hp, maxHp, sp, maxSp Sp, hunger Hunger)
	GetStatsString() string
}
