package effect

//go:generate stringer -type=Effect
type Effect int

const (
	Starving Effect = iota
	Poisoned
	Bleeding
)

var EffectList = [...]Effect{
	Starving,
	Bleeding,
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
	case Starving, Poisoned, Bleeding:
		return Debuff
	}
	return Other
}
