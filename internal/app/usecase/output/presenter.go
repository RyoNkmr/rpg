package output

import (
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity"
	"github.com/rivo/tview"
)

type RootPresenter interface {
	Run() error
}

type StatusPresenter interface {
	GetView() *tview.Table
	Update()
}

type MainPresenter interface {
	GetView() *tview.TextView
	Hr()
	AddLine(string)
	AddLines(messages []string, delay time.Duration)
	AddPositiveLine(string)
	AddNegativeLine(string)
}

type InventoryPresenter interface {
	GetView() *tview.Box
}

type CommandPresenter interface {
	GetView() *tview.List
	Clear()
	WaitFor([]*entity.Command)
}
