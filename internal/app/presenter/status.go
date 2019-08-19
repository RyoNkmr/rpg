package presenter

import (
	"fmt"
	"math"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type statusPresenter struct {
	table  *tview.Table
	player player.Player
}

func NewStatusPresenter(player player.Player) *statusPresenter {
	table := tview.NewTable().
		SetBorders(true)
	return &statusPresenter{table, player}
}

func (p *statusPresenter) GetView() *tview.Table {
	return p.table
}

func (p *statusPresenter) Update() {
	level := p.player.GetCurrentLevel()
	exp := p.player.GetExpToNextLevel()
	hp, hpmax, sp, spmax, hunger := p.player.GetStats()

	data := [...][2]string{
		{"Level", fmt.Sprintf("%d", level)},
		{"Next Level", fmt.Sprintf("%d", exp)},
		{"Hp", fmt.Sprintf("%d / %d", hp, hpmax)},
		{"Sp", fmt.Sprintf("%d / %d", sp, spmax)},
		{"Hunger", fmt.Sprintf("%d", hunger)},
	}

	colors := [...][2]tcell.Color{
		{tcell.ColorWhite, tcell.ColorWhite},
		{tcell.ColorWhite, tcell.ColorWhite},
		{tcell.ColorWhite, getPercentColor(hp, hpmax)},
		{tcell.ColorWhite, getPercentColor(sp, spmax)},
		{tcell.ColorWhite, getPercentColor(hunger, 100)},
	}

	for col, colData := range data {
		for row, text := range colData {
			cell := tview.NewTableCell(text).
				SetExpansion(1).
				SetTextColor(colors[col][row]).
				SetAlign(tview.AlignCenter)
			p.table.SetCell(row, col, cell)
		}
	}
}

func getPercentColor(dividend int64, divisor int64) tcell.Color {
	if divisor == 0 {
		return tcell.ColorWhite
	}
	switch v := math.Floor(float64(dividend) / float64(divisor) * 100); {
	case v >= 75:
		return tcell.ColorGreen
	case v >= 50:
		return tcell.ColorYellow
	case v >= 30:
		return tcell.ColorOrange
	default:
		return tcell.ColorRed
	}
}
