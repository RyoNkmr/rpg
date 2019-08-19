package playerclass

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
)

type PlayerClass interface {
	fmt.Stringer
	GetHitDiceBonus() dice.DiceSide
}
