package presenter

import "github.com/rivo/tview"

type commandPresenter struct {
	view *tview.Box
}

func NewCommandPresenter() *commandPresenter {
	view := tview.NewBox().
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("commands")

	return &commandPresenter{view}
}

func (p *commandPresenter) GetView() *tview.Box {
	return p.view
}
