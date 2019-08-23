package usecase

import (
	"fmt"
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/usecase/output"
)

type battleUsecase struct {
	main      output.MainPresenter
	status    output.StatusPresenter
	inventory output.InventoryPresenter
}

type BattleUsecase interface {
	HandleAttack(attacker, receiver actor.Actor, deathMes string) (isReceiverDead bool, isAttackerDead bool)
}

func NewBattleUsecase(main output.MainPresenter, status output.StatusPresenter, inventory output.InventoryPresenter) *battleUsecase {
	return &battleUsecase{main, status, inventory}
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
