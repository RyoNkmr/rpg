package playerclass

type warrior struct{}

func NewWarrior() *warrior {
	return &warrior{}
}

func (p *warrior) String() string {
	return "warriror"
}
