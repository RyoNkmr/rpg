package player

import (
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/class"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player/race"
)

type player struct {
	hp     int64
	hpmax  int64
	mp     int64
	mpmax  int64
	hunger int64
	// skills []*skill.Skill
}

type Player interface {
	race.Race
	class.Class
}
