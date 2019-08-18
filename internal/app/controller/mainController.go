package controller

import (
	"fmt"
	"strings"
	"time"

	"github.com/RyoNkmr/rpg/internal/app/entity/actor"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/enemy"
	"github.com/RyoNkmr/rpg/internal/app/entity/actor/player"
	"github.com/rivo/tview"
)

type mainController struct {
	App        *tview.Application
	rootView   *tview.Flex
	status     *tview.Box
	mainscreen *tview.TextView
	commands   *tview.Box
	items      *tview.Box
	Player     player.Player
}

type MainController interface {
	Run() error
}

func NewMainController(App *tview.Application, Player player.Player) *mainController {
	// status := tview.NewTable().
	// 	SetBorders(true).
	status := tview.NewBox().
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("status")

	mainscreen := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			App.Draw()
		})

	commands := tview.NewBox().
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("commands")

	items := tview.NewBox().
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("items")

	leftColumn := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(status, 5, 1, false).
		AddItem(mainscreen, 0, 3, false).
		AddItem(commands, 8, 2, false)
	rightColumn := items

	rootView := tview.NewFlex().
		AddItem(leftColumn, 0, 1, false).
		AddItem(rightColumn, 20, 1, false)

	App.SetRoot(rootView, true)

	return &mainController{
		App,
		rootView,
		status,
		mainscreen,
		commands,
		items,
		Player,
	}
}

func (c *mainController) drawText(s string) {
	ss := strings.Split(s, " ")
	lastIndex := len(ss) - 1

	for i, word := range ss {
		if word == "damage" {
			continue
		}
		if i != lastIndex && ss[i+1] == "damage" {
			word = "[red]" + word + " damage[white]"
		}
		fmt.Fprint(c.mainscreen, word+" ")
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Fprint(c.mainscreen, "\n")
}

func (c *mainController) drawline() {
	fmt.Fprintln(c.mainscreen, "----------------------")
}

func (c *mainController) start() {
	enemy := enemy.NewSnake()

	c.drawline()
	c.drawText(c.Player.GetStats())
	c.drawText(enemy.GetStats())
	c.drawline()

	for {
		time.Sleep(600 * time.Millisecond)
		if c.handleTurn(c.Player, enemy, "won") {
			break
		}
		time.Sleep(600 * time.Millisecond)
		if c.handleTurn(enemy, c.Player, "gameover") {
			break
		}
	}
}

func (c *mainController) Run() error {
	go c.start()
	if err := c.App.Run(); err != nil {
		return err
	}
	defer c.App.Stop()
	return nil
}

func (c *mainController) dumpMessages(ms []actor.Message) {
	for _, m := range ms {
		time.Sleep(600 * time.Millisecond)
		c.drawText(m)
	}
}

func (c *mainController) handleTurn(attacker, receiver actor.Actor, deadMes string) bool {
	pad, ams := attacker.Attack(receiver)
	dms, isReceiverDead := receiver.Damage(pad)
	c.dumpMessages(append(ams, dms))
	c.dumpBetweenLines(receiver.GetStats())
	if isReceiverDead {
		c.drawText(deadMes)
	}
	return isReceiverDead
}

func (c *mainController) dumpBetweenLines(m string) {
	time.Sleep(600 * time.Millisecond)
	c.drawline()
	c.drawText(m)
	c.drawline()
}
