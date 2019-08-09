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

	stats := &race{
		Strength:     strd.Cast(),
		Intelligence: intd.Cast(),
		Dexterity:    dexd.Cast(),
		Constitution: cond.Cast(),
		ExpRate:      1.2,
	}
	return &dwarf{stats}
}

func (r *dwarf) String() string {
	return "dwarf"
}
