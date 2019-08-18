package player

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/playerclass"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/race"
	"github.com/RyoNkmr/rpg/internal/app/entity/dice"
)

type player struct {
	Hp     int64
	Hpmax  int64
	Sp     int64
	Spmax  int64
	Hunger int64
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
		Hp:     10,
		Hpmax:  10,
		Sp:     10,
		Spmax:  10,
		Hunger: 100,
		race:   race,
		class:  class,
	}
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
	p.Hp -= int64(d)
	return message, p.Hp <= 0
}

func (p *player) GetName() string {
	return fmt.Sprintf("player, the %s %s", p.race, p.class)
}

func (p *player) GetStats() string {
	return fmt.Sprintf("player: Hp: %d", p.Hp)
}
