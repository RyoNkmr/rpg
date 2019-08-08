package race

type human struct {
	*race
}

func NewHuman() *human {
	stats := &race{
		Strength:     8,
		Intelligence: 8,
		Dexterity:    8,
		Constitution: 8,
		ExpRate:      0.9,
	}
	return &human{stats}
}
