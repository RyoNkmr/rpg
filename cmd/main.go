package main

import (
	"fmt"
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/playerclass"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/race"
)

func main() {
	enemy := enemy.NewSnake()

	race := race.NewHuman()
	class := playerclass.NewWarrior()
	player := player.NewPlayer(race, class)

	line()
	fmt.Println(player.GetStats())
	fmt.Println(enemy.GetStats())
	line()

	for {
		time.Sleep(600 * time.Millisecond)
		if handleTurn(player, enemy, "won") {
			break
		}
		time.Sleep(600 * time.Millisecond)
		if handleTurn(enemy, player, "gameover") {
			break
		}
	}
}

func dumpMessages(ms []actor.Message) {
	for _, m := range ms {
		time.Sleep(600 * time.Millisecond)
		fmt.Println(m)
	}
}

func handleTurn(attacker, receiver actor.Actor, deadMes string) bool {
	pad, ams := attacker.Attack(receiver)
	dms, isReceiverDead := receiver.Damage(pad)
	dumpMessages(append(ams, dms))
	lines(receiver.GetStats())
	if isReceiverDead {
		fmt.Println(deadMes)
	}
	return isReceiverDead
}

func line() {
	fmt.Println("----------------------")
}

func lines(m string) {
	time.Sleep(600 * time.Millisecond)
	fmt.Println("----------------------")
	fmt.Println(m)
	fmt.Println("----------------------")
}
