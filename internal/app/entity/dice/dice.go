package dice

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/RyoNkmr/rpg/pkg"
)

type DiceValue = uint64
type DiceNumber uint16
type DicePer uint16
type DiceSide uint16

type dice struct {
	number DiceNumber
	side   DiceSide
	rnd    *rand.Rand
}

type Dice interface {
	fmt.Stringer
	Cast() DiceValue
	PCast() (DiceValue, DicePer)
	GetMax() DiceValue
	GetMix() DiceValue
}

func NewDice(number DiceNumber, side DiceSide) *dice {
	rnd := pkg.GetRand()
	return &dice{number, side, rnd}
}

func (d *dice) castOne() DiceValue {
	return DiceValue(d.rnd.Int63n(int64(d.side)) + 1)
}

func (d *dice) Cast() (v DiceValue) {
	cnt := d.number
	for cnt > 0 {
		v = v + d.castOne()
		cnt--
	}
	return v
}

func getPercent(value DiceValue, max DiceValue) DicePer {
	return DicePer(math.Floor(float64(value) / float64(max) * 100))
}

func (d *dice) PCast() (v DiceValue, percent DicePer) {
	v = d.Cast()
	return v, getPercent(v, d.GetMax())
}

func (d *dice) GetMax() DiceValue {
	return DiceValue(d.number) * DiceValue(d.side)
}

func (d *dice) GetMix() DiceValue {
	return DiceValue(d.number)
}

func (d *dice) String() string {
	return fmt.Sprintf("%dd%d", d.number, d.side)
}
