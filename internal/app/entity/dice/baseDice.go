package dice

import (
	"fmt"
)

type baseValueDice struct {
	*dice
	base DiceValue
}

func NewBaseValueDice(number DiceNumber, side DiceSide, base DiceValue) *baseValueDice {
	dice := NewDice(number, side)
	return &baseValueDice{dice, base}
}

func (d *baseValueDice) Cast() (v DiceValue) {
	return d.dice.Cast() + d.base
}

func (d *baseValueDice) PCast() (v DiceValue, percent DicePer) {
	v = d.Cast()
	return v, getPercent(v, d.GetMax())
}

func (d *baseValueDice) GetMax() DiceValue {
	return DiceValue(d.number)*DiceValue(d.side) + d.base
}

func (d *baseValueDice) GetMix() DiceValue {
	return DiceValue(d.number) + d.base
}

func (d *baseValueDice) String() string {
	return fmt.Sprintf("%dd%d+%d", d.number, d.side, d.base)
}
