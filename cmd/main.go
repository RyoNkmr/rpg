package main

import (
	"github.com/RyoNkmr/rpg/internal/app/controller"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/playerclass"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/race"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	race := race.NewHuman()
	class := playerclass.NewWarrior()
	player := player.NewPlayer(race, class)

	c := controller.NewMainController(app, player)

	if err := c.Run(); err != nil {
		panic(err)
	}

}
