package player

import (
	"fmt"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/effect"
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

	hitDice       dice.Dice
	gainedHpTable []actor.Hp
	effects       effect.EffectMap
}

type Player interface {
	actor.Actor
	GainExp(e actor.Exp) (isLevelChanged bool)
	GetCurrentLevel() actor.Level
	GetExpToNextLevel() actor.Exp
	GetExp() actor.Exp
}

func NewPlayer(race race.Race, class playerclass.PlayerClass) *player {
	hitDice := dice.NewDice(1, race.GetHitDiceBase()+class.GetHitDiceBonus())
	return &player{
		hp:      10,
		hpmax:   10,
		sp:      10,
		spmax:   10,
		hunger:  100,
		race:    race,
		class:   class,
		hitDice: hitDice,
		effects: effect.EffectMap{},
	}
}

func (p *player) IsFriend() bool {
	return true
}

func (p *player) IsAlive() bool {
	return p.hp > 0
}

func (p *player) metabolize() {
	if p.hunger > 0 {
		p.hunger -= 2
	}
	if p.hunger <= 0 {
		p.effects.Add(effect.Starving)
	}
}

func (p *player) GetEffects() []effect.Effect {
	return p.effects.AsOrderedList()
}

func (p *player) AddEffect(e effect.Effect) {
	p.effects.Add(e)
}

func (p *player) RemoveEffect(e effect.Effect) {
	p.effects.Remove(e)
}

func (p *player) handleDamageEffect(messages *[]actor.Message, e effect.Effect, d actor.Damage, m actor.Message) (isDead bool) {
	if !p.effects.Has(e) {
		return false
	}
	*messages = append(*messages, m)
	dm, isDead := p.Damage(d)
	*messages = append(*messages, dm)
	return isDead
}

func (p *player) handleEffects() (messages []actor.Message, isDead bool) {
	if p.handleDamageEffect(&messages, effect.Starving, actor.Damage(p.hp/10), "you are starving!") {
		return messages, true
	}
	if p.handleDamageEffect(&messages, effect.Poisoned, actor.Damage(p.hpmax/10), "you are poisoned!") {
		return messages, true
	}
	// NOTE: DRYじゃないけど、今後実装が多様になるので
	return messages, isDead
}

func (p *player) BeforeAttack() (messages []actor.Message, isDead bool) {
	p.metabolize()
	return p.handleEffects()
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
	return fmt.Sprintf("player the %s %s", p.race, p.class)
}

func (p *player) String() string {
	return p.GetName()
}

func (p *player) GetStatsString() string {
	return fmt.Sprintf("player: hp: %d", p.hp)
}

func (p *player) GetStats() (hp, maxHp actor.Hp, sp, maxSp actor.Sp, hunger actor.Hunger) {
	return p.hp, p.hpmax, p.sp, p.spmax, p.hunger
}

func (p *player) GainExp(e actor.Exp) (isLevelChanged bool) {
	isLevelChanged = p.race.GainExp(e)
	if isLevelChanged {
		p.onLevelChanged(true)
	}
	return isLevelChanged
}

func (p *player) GetCurrentLevel() actor.Level {
	return p.race.GetCurrentLevel()
}

func (p *player) onLevelChanged(isUpper bool) {
	if !isUpper {
		lastIndex := len(p.gainedHpTable) - 1
		lostHp := p.gainedHpTable[lastIndex]
		p.gainedHpTable = p.gainedHpTable[:lastIndex]
		p.hpmax -= lostHp
		if p.hp > p.hpmax {
			p.hp = p.hpmax
		}
		return
	}

	hpGained := actor.Hp(p.hitDice.GetMax())
	if p.GetCurrentLevel() > 2 {
		hpGained = actor.Hp(p.hitDice.Cast())
	}

	p.gainedHpTable = append(p.gainedHpTable, hpGained)
	p.hpmax += hpGained
	p.hp = p.hpmax
}

func (p *player) GetHpBonus() actor.Hp {
	v := actor.Hp(0)
	for _, hp := range p.gainedHpTable {
		v += hp
	}
	return v
}

func (p *player) GetExp() actor.Exp {
	return p.race.GetExp()
}

// func (p *player) LoseExp(e actor.Exp) (isLevelChanged bool) {
// 	return p.race.LoseExp(e)
// }
func (p *player) GetExpToNextLevel() actor.Exp {
	return p.race.GetExpToNextLevel()
}

// func (p *player) GetExpToLevel(l actor.Level) actor.Exp {
// 	return p.race.GetExpToLevel(l)
// }
