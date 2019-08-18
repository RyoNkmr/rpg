package presenter

import (
	"github.com/RyoNkmr/rpg/internal/app/usecase/output"
	"github.com/google/wire"
	"github.com/rivo/tview"
)

var PresenterSet = wire.NewSet(
	tview.NewApplication,
	NewCommandPresenter,
	NewInventoryPresenter,
	NewMainPresenter,
	NewRootPresenter,
	NewStatusPresenter,
	wire.Bind(new(output.CommandPresenter), new(*commandPresenter)),
	wire.Bind(new(output.InventoryPresenter), new(*inventoryPresenter)),
	wire.Bind(new(output.MainPresenter), new(*mainPresenter)),
	wire.Bind(new(output.RootPresenter), new(*rootPresenter)),
	wire.Bind(new(output.StatusPresenter), new(*statusPresenter)),
)
