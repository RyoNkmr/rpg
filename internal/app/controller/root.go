package controller

import (
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
	"github.com/RyoNkmr/rpg/internal/app/usecase"
)

type rootController struct {
	player player.Player
	battle usecase.BattleUsecase
	system usecase.SystemUsecase
	d      dice.Dice
}

type RootController interface {
	Run() error
}

func NewRootController(player player.Player, battle usecase.BattleUsecase, system usecase.SystemUsecase) *rootController {
	d := dice.NewDice(1, 3)
	return &rootController{
		player,
		battle,
		system,
		d,
	}
}

func (c *rootController) Run() error {
	go c.run()
	return c.system.Run()
}

func (c *rootController) run() {
	wcount := 0
	for c.player.IsAlive() {
		enemy := c.walk(wcount)
		if enemy == nil {
			wcount++
			continue
		}

		wcount = 0
		c.system.AddLine("you encounter " + enemy.GetName())
		c.sleep()
		c.system.AddLineBetweenHr(enemy.GetStatsString())

		c.handleBattle(enemy)
	}
}

func (c *rootController) handleBattle(enemy enemy.Enemy) {
	for {
		c.sleep()
		if c.battle.HandleAttack(c.player, enemy, "won") {
			return
		}
		c.sleep()
		if c.battle.HandleAttack(enemy, c.player, "gameover") {
			return
		}
	}
}

func (c *rootController) walk(wcount int) enemy.Enemy {
	str := "you go"
	for i := 0; i < wcount; i++ {
		str += " on and on"
	}
	c.system.AddLine(str + " this road")

	if c.d.Cast() == 3 {
		return c.encounter()
	}

	return nil
}

func (c *rootController) encounter() enemy.Enemy {
	enemy := enemy.NewSnake()
	return enemy
}

func (c *rootController) sleep() {
	time.Sleep(600 * time.Millisecond)
}
