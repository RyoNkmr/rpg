package entity

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/RyoNkmr/rpg/pkg"
)

type Dice struct {
	number uint16
	side   uint16
	rnd    *rand.Rand
}

func NewDice(number uint16, side uint16) *Dice {
	rnd := pkg.GetRand()
	return &Dice{number, side, rnd}
}

func (d *Dice) castOne() uint64 {
	return uint64(d.rnd.Int63n(int64(d.side)) + 1)
}

func (d *Dice) Cast() (v uint64) {
	cnt := d.number
	for cnt > 0 {
		v = v + d.castOne()
		cnt--
	}
	return v
}

func (d *Dice) PCast() (v uint64, percent uint64) {
	v = d.Cast()
	return v, d.GetPercent(v)
}

func (d *Dice) GetPercent(v uint64) uint64 {
	return uint64(math.Floor(float64(v) / float64(d.GetMax()) * 100))
}

func (d *Dice) GetMax() uint64 {
	return uint64(d.number) * uint64(d.side)
}

func (d *Dice) String() string {
	return fmt.Sprintf("%dd%d", d.number, d.side)
}
