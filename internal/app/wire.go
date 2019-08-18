// +build wireinject

package app

import (
	"github.com/RyoNkmr/rpg/internal/app/controller"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/presenter"
	"github.com/RyoNkmr/rpg/internal/app/usecase"
	"github.com/google/wire"
)

func Inject(player player.Player) (controller.RootController, error) {
	wire.Build(
		usecase.UsecaseSet,
		presenter.PresenterSet,
		controller.ControllerSet,
	)
	return nil, nil
}
