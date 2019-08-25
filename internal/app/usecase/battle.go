package usecase

import (
	"fmt"
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/usecase/output"
)

type battleUsecase struct {
	main      output.MainPresenter
	status    output.StatusPresenter
	inventory output.InventoryPresenter
	command   output.CommandPresenter
}

type BattleUsecase interface {
	WaitForCommands(player player.Player, enemy enemy.Enemy) (isPlayerDead, isEnemyDead bool)
	HandleAttack(attacker, receiver actor.Actor, deathMes string) (isReceiverDead bool, isAttackerDead bool)
	HandlePray(player player.Player)
}

func NewBattleUsecase(main output.MainPresenter, status output.StatusPresenter, inventory output.InventoryPresenter, command output.CommandPresenter) *battleUsecase {
	return &battleUsecase{main, status, inventory, command}
}

func (u *battleUsecase) WaitForCommands(player player.Player, enemy enemy.Enemy) (isPlayerDead, isEnemyDead bool) {
	ch := make(chan bool, 2)
	atk := &entity.Command{
		Text:        "Attack",
		ShortCutKey: 'a',
		Callback: func() {
			isEnemyDead, isPlayerDead := u.HandleAttack(player, enemy, "won")
			ch <- isEnemyDead
			ch <- isPlayerDead
		},
	}

	pry := &entity.Command{
		Text:        "Prey",
		ShortCutKey: 'p',
		Callback: func() {
			u.HandlePray(player)
			ch <- false
			ch <- false
		},
	}

	commands := []*entity.Command{atk, pry}
	u.command.WaitFor(commands)
	isPlayerDead, isEnemyDead = <-ch, <-ch
	return isPlayerDead, isEnemyDead
}

func (u *battleUsecase) HandlePray(player player.Player) {
	u.main.AddLine(player.GetName() + " prayed for yourself.")
}

func (u *battleUsecase) HandleAttack(attacker, receiver actor.Actor, deathMes string) (isReceiverDead bool, isAttackerDead bool) {
	bms, isAttackerDead := attacker.BeforeAttack()
	u.main.AddLines(bms, 600*time.Millisecond)
	u.status.Update()
	if isAttackerDead {
		return false, isAttackerDead
	}

	pad, ams := attacker.Attack(receiver)
	dms, isReceiverDead := receiver.Damage(pad)

	isReceiverEnemy := !receiver.IsFriend()

	u.main.AddLines(append(ams, dms), 600*time.Millisecond)
	u.status.Update()

	u.main.Hr()
	u.handleConditionalMessage(!isReceiverEnemy, receiver.GetStatsString())
	u.main.Hr()

	if isReceiverDead {
		u.handleConditionalMessage(isReceiverEnemy, deathMes)

		if isReceiverEnemy {
			exp := receiver.(enemy.Enemy).GetExp()
			player := attacker.(player.Player)
			u.main.AddPositiveLine(fmt.Sprintf("you gained %d exp", exp))

			if levelUp := player.GainExp(exp); levelUp {
				level := player.GetCurrentLevel()
				u.main.AddPositiveLine(fmt.Sprintf("welcome you have been reached to level %d", level))
			} else {
				exp := player.GetExpToNextLevel()
				u.main.AddLine(fmt.Sprintf("%d more exp to the next level", exp))
			}

			u.status.Update()
		}
	}

	return isReceiverDead, false
}

func (u *battleUsecase) handleConditionalMessage(isPositive bool, mes string) {
	if isPositive {
		u.main.AddPositiveLine(mes)
		return
	}
	u.main.AddPositiveLine(mes)
}
