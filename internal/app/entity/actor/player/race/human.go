package race

type human struct {
	*race
}

func NewHuman() *human {
	stats := newRace(
		8,
		8,
		8,
		8,
		0.9,
		10,
	)
	return &human{stats}
}

func (r *human) String() string {
	return "human"
}
