package presenter

import (
	"github.com/RyoNkmr/rpg/internal/app/usecase/output"
	"github.com/rivo/tview"
)

type rootPresenter struct {
	app *tview.Application
}

func NewRootPresenter(app *tview.Application, s output.StatusPresenter, m output.MainPresenter, c output.CommandPresenter, i output.InventoryPresenter) *rootPresenter {

	status := s.GetView()
	mainscreen := m.GetView()
	commands := c.GetView()
	inventory := i.GetView()

	leftColumn := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(status, 0, 1, false).
		AddItem(mainscreen, 0, 2, false).
		AddItem(commands, 8, 1, false)
	rightColumn := inventory

	rootView := tview.NewFlex().
		AddItem(leftColumn, 0, 1, false).
		AddItem(rightColumn, 20, 1, false)

	app.SetRoot(rootView, true)

	return &rootPresenter{app}
}

func (p *rootPresenter) Run() error {
	if err := p.app.Run(); err != nil {
		return err
	}
	defer p.app.Stop()
	return nil
}
