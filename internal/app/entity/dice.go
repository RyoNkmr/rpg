package entity

import (
	"math/rand"

	"github.com/RyoNkmr/rpg/pkg"
)

type Dice struct {
	number uint16
	side   uint16
	rnd    *rand.Rand
}

func NewDice(number uint16, dside uint16) *Dice {
	side := dside + 1
	rnd := pkg.GetRand()
	return &Dice{number, side, rnd}
}

func (d *Dice) castOne() uint32 {
	return uint32(d.rnd.Int31n(int32(d.side)))
}

func (d *Dice) Cast() (v uint64) {
	cnt := d.number
	for cnt > 0 {
		v = v + uint64(d.castOne())
		cnt--
	}
	return v
}
