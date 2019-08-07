package main

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity"
)

func main() {
	dice := entity.NewDice(12, 19)
	fmt.Println(dice.String())
	fmt.Println(dice.GetMax())
	fmt.Println(dice.Cast())
	fmt.Println(dice.Cast())
	fmt.Println(dice.PCast())
	fmt.Println(dice.PCast())
	fmt.Println(dice.PCast())
}
