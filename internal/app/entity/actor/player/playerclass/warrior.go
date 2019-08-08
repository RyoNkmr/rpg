package playerclass

import (
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
)

type warrior struct{}

func (p *warrior) New() *warriror {
	return &warriror{}
}

func (p *warrior) Attack(e *enemy.Enemy) {
}
