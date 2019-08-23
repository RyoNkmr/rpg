package effect

//go:generate stringer -type=Effect
type Effect int

const (
	Starving Effect = iota
	Poisoned
)

var EffectList = [...]Effect{
	Starving,
	Poisoned,
}

//------------

type EffectType int

const (
	Buff EffectType = iota
	Debuff
	Other
)

func (e Effect) GetType() EffectType {
	switch e {
	case Starving, Poisoned:
		return Debuff
	}
	return Other
}
