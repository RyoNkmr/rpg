package main

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity"
)

func main() {
	dice := entity.NewDice(4, 5)
	fmt.Println(dice.Cast())
	fmt.Println(dice.Cast())
	fmt.Println(dice.Cast())
	fmt.Println(dice.Cast())
	fmt.Println(dice.Cast())
}
