package main

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/playerclass"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/race"
)

func DumpMessages(ms []actor.Message) {
	for _, m := range ms {
		fmt.Println(m)
	}
}

func main() {
	enemy := enemy.NewSnake()

	race := race.NewHuman()
	class := playerclass.NewWarrior()
	player := player.NewPlayer(race, class)

	fmt.Println(player.GetStats())
	fmt.Println(enemy.GetStats())

	for i := 0; i < 20; i++ {
		pad, ams := player.Attack(enemy)
		dms, isEnemyDead := enemy.Damage(pad)
		DumpMessages(append(ams, dms))
		fmt.Println(enemy.GetStats())
		if isEnemyDead {
			fmt.Println("clear")
			break
		}

		ead, ams := enemy.Attack(player)
		dms, isPlayerDead := player.Damage(ead)
		DumpMessages(append(ams, dms))
		fmt.Println(player.GetStats())
		if isPlayerDead {
			fmt.Println("game over")
			break
		}
	}
}
