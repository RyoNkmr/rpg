package presenter

import (
	"fmt"
	"strings"
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/rivo/tview"
)

type mainPresenter struct {
	view *tview.TextView
}

func NewMainPresenter(app *tview.Application) *mainPresenter {
	view := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	return &mainPresenter{view}
}

func (p *mainPresenter) GetView() *tview.TextView {
	return p.view
}

func (p *mainPresenter) AddLine(s string) {
	ss := strings.Split(s, " ")
	lastIndex := len(ss) - 1

	for i, word := range ss {
		if word == "damage" {
			continue
		}
		if i != lastIndex && ss[i+1] == "damage" {
			word = "[red]" + word + " damage[-]"
		}
		fmt.Fprint(p.view, word+" ")
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Fprint(p.view, "\n")
}

func (p *mainPresenter) AddNegativeLine(s string) {
	p.addColoredLine(s, "red")
}

func (p *mainPresenter) AddPositiveLine(s string) {
	p.addColoredLine(s, "green")
}

func (p *mainPresenter) AddLines(ms []actor.Message, delay time.Duration) {
	for _, m := range ms {
		p.AddLine(m)
		time.Sleep(delay)
	}
}

func (p *mainPresenter) Hr() {
	fmt.Fprintln(p.view, "---")
}

func (p *mainPresenter) addColoredLine(s string, color string) {
	ss := strings.Split(s, " ")
	fmt.Fprint(p.view, "["+color+"]")
	for _, word := range ss {
		fmt.Fprint(p.view, word+" ")
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Fprint(p.view, "[-]\n")
}
