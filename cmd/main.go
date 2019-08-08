package main

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
)

func main() {
	dice := dice.NewDice(12, 19)
	fmt.Println(dice)
	fmt.Println(dice.GetMax())
	fmt.Println(dice.GetMix())
	fmt.Println(dice.Cast())
	fmt.Println(dice.Cast())
	fmt.Println(dice.PCast())
	fmt.Println(dice.PCast())
	fmt.Println(dice.PCast())
}
