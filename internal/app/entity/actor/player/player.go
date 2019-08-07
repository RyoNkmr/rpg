package player

type player struct {
	hp     uint64
	hpmax  uint64
	mp     uint64
	mpmax  uint64
	hunger uint64
	items  []*Item
}

func NewPlayer(hp uint64, mp uint64) *player {
	return &player{hp, mp}
}

type Player interface {
	Attack(*Enemy)
	ConfusedAttack(*Player)
	PowerAttack(*Enemy)
	SpecialAttack(*Enemy)
}
