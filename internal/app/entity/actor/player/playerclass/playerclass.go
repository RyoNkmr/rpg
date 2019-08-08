package playerclass

import (
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
	"github.com/RyoNkmr/rpg/internal/app/entity/item"
)

type playerClass struct {
	items []*item.Item
}

type PlayerClass interface {
	Attack(*enemy.Enemy)
	Throw(*item.Item, *enemy.Enemy)
}
