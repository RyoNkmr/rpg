package presenter

import (
	"github.com/RyoNkmr/rpg/internal/app/entity"
	"github.com/rivo/tview"
)

type commandPresenter struct {
	view *tview.List
	app  *tview.Application
}

func NewCommandPresenter(app *tview.Application) *commandPresenter {
	view := tview.NewList()
	view.SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("commands")

	return &commandPresenter{view, app}
}

func (p *commandPresenter) GetView() *tview.List {
	return p.view
}

func (p *commandPresenter) Clear() {
	p.view.Clear()
}

func (p *commandPresenter) WaitFor(commands []*entity.Command) {
	p.Clear()
	for _, c := range commands {
		fn := c.Callback
		p.view.AddItem(c.Text, c.SecondaryText, c.ShortCutKey, func() {
			go fn()
			p.Clear()
		})
	}
	p.app.Draw().SetFocus(p.view)
}
