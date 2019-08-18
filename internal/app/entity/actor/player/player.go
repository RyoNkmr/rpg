package player

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/playerclass"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/race"
	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
)

type player struct {
	hp     actor.Hp
	hpmax  actor.Hp
	sp     actor.Sp
	spmax  actor.Sp
	hunger actor.Hunger
	// skills []*skill.Skill
	race  race.Race
	class playerclass.PlayerClass
}

type Player interface {
	actor.Actor
	// Damage(actor.Damage) (message actor.Message, isDead bool)
}

func NewPlayer(race race.Race, class playerclass.PlayerClass) *player {
	return &player{
		hp:     10,
		hpmax:  10,
		sp:     10,
		spmax:  10,
		hunger: 100,
		race:   race,
		class:  class,
	}
}

func (p *player) IsFriend() bool {
	return true
}

func (p *player) IsAlive() bool {
	return p.hp > 0
}

func (p *player) Attack(t actor.Actor) (actor.Damage, []actor.Message) {
	dice := p.GetAttackDice()
	m := make([]actor.Message, 0)
	m = append(m, fmt.Sprintf("you attack %s", t.GetName()))
	return dice.Cast(), m
}

func (p *player) GetAttackDice() dice.Dice {
	return dice.NewDice(2, 2)
}

func (p *player) Damage(d actor.Damage) (message actor.Message, isDead bool) {
	message = fmt.Sprintf("you take %d damage", d)
	p.hp -= int64(d)
	return message, p.hp <= 0
}

func (p *player) GetName() string {
	return fmt.Sprintf("player, the %s %s", p.race, p.class)
}

func (p *player) GetStatsString() string {
	return fmt.Sprintf("player: hp: %d", p.hp)
}

func (p *player) GetStats() (hp, maxHp actor.Hp, sp, maxSp actor.Sp, hunger actor.Hunger) {
	return p.hp, p.hpmax, p.sp, p.spmax, p.hunger
}
