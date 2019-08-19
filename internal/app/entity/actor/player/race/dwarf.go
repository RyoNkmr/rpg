package race

import "github.com/RyoNkmr/rpg/internal/app/entity/dice"

type dwarf struct {
	*race
}

func NewDwarf() *dwarf {
	strd := dice.NewBaseValueDice(2, 3, 9)
	intd := dice.NewBaseValueDice(1, 4, 3)
	dexd := dice.NewBaseValueDice(2, 4, 5)
	cond := dice.NewBaseValueDice(2, 3, 8)

	stats := newRace(
		strd.Cast(),
		intd.Cast(),
		dexd.Cast(),
		cond.Cast(),
		1.2,
		11,
	)
	return &dwarf{stats}
}

func (r *dwarf) String() string {
	return "dwarf"
}
