package usecase

import (
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/usecase/output"
)

type attackUsecase struct {
	main      output.MainPresenter
	status    output.StatusPresenter
	inventory output.InventoryPresenter
}

type AttackUsecase interface {
	HandleAttack(attacker, receiver actor.Actor, deathMes string) (isReceiverDead bool)
}

func NewAttackUsecase(main output.MainPresenter, status output.StatusPresenter, inventory output.InventoryPresenter) *attackUsecase {
	return &attackUsecase{main, status, inventory}
}

func (u *attackUsecase) HandleAttack(attacker, receiver actor.Actor, deathMes string) (isReceiverDead bool) {
	pad, ams := attacker.Attack(receiver)
	dms, isReceiverDead := receiver.Damage(pad)

	u.main.AddLines(append(ams, dms), 600*time.Millisecond)
	u.status.Update()

	u.main.Hr()
	u.handleConditionalMessage(receiver.IsFriend(), receiver.GetStatsString())
	u.main.Hr()

	if isReceiverDead {
		u.handleConditionalMessage(!receiver.IsFriend(), deathMes)
	}

	return isReceiverDead
}

func (u *attackUsecase) handleConditionalMessage(isPositive bool, mes string) {
	if isPositive {
		u.main.AddPositiveLine(mes)
		return
	}
	u.main.AddPositiveLine(mes)
}
