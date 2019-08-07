package enemy

import "github.com/RyoNkmr/rpg/internal/app/entity/actor/player"

type enemy struct {
	hp uint64
	mp uint64
}

func NewEnemy(hp uint64, mp uint64) *enemy {
	return &enemy{hp, mp}
}

type Enemy interface {
	Attack(*player.Player)
	ConfusedAttack(*Enemy)
	SpecialAttack(*player.Player)
}
