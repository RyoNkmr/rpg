package presenter

import "github.com/rivo/tview"

type inventoryPresenter struct {
	view *tview.Box
}

func NewInventoryPresenter() *inventoryPresenter {
	view := tview.NewBox().
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("items")

	return &inventoryPresenter{view}
}

func (p *inventoryPresenter) GetView() *tview.Box {
	return p.view
}
