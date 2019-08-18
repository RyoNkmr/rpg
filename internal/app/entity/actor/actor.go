package actor

type Message = string
type Damage = uint64

type Hp = int64
type Sp = int64
type Hunger = int64

type Actor interface {
	Attack(Actor) (Damage, []Message)
	Damage(Damage) (message Message, isDead bool)
	IsFriend() bool
	IsAlive() bool
	GetName() string
	GetStats() (hp, maxHp, sp, maxSp Sp, hunger Hunger)
	GetStatsString() string
}
