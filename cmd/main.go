package main

import (
	"github.com/RyoNkmr/rpg/internal/app"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/playerclass"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/race"
)

func main() {
	race := race.NewHuman()
	class := playerclass.NewWarrior()
	player := player.NewPlayer(race, class)

	ctrl, err := app.Inject(player)

	if err != nil {
		panic(err)
	}

	if err := ctrl.Run(); err != nil {
		panic(err)
	}

}
