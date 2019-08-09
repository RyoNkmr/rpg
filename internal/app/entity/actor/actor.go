package actor

type Message = string
type Damage = uint64

type Actor interface {
	Attack(Actor) (Damage, []Message)
	Damage(Damage) (message Message, isDead bool)
	GetName() string
	GetStats() string
}
