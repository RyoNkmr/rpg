// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"github.com/RyoNkmr/rpg/internal/app/controller"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/RyoNkmr/rpg/internal/app/presenter"
	"github.com/RyoNkmr/rpg/internal/app/usecase"
	"github.com/rivo/tview"
)

// Injectors from wire.go:

func Inject(player2 player.Player) (controller.RootController, error) {
	application := tview.NewApplication()
	mainPresenter := presenter.NewMainPresenter(application)
	statusPresenter := presenter.NewStatusPresenter(player2)
	inventoryPresenter := presenter.NewInventoryPresenter()
	battleUsecase := usecase.NewBattleUsecase(mainPresenter, statusPresenter, inventoryPresenter)
	commandPresenter := presenter.NewCommandPresenter()
	rootPresenter := presenter.NewRootPresenter(application, statusPresenter, mainPresenter, commandPresenter, inventoryPresenter)
	systemUsecase := usecase.NewSystemUsecase(rootPresenter, statusPresenter, mainPresenter)
	rootController := controller.NewRootController(player2, battleUsecase, systemUsecase)
	return rootController, nil
}
