package show2b

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zengqiang96/grab2b/internal/api"
)

type TUI struct {
	client     *api.Client
	app        *tview.Application
	layout     *tview.Flex
	leftPanel  *tview.Flex
	rightPanel *tview.Flex

	yearPanel *tview.InputField
	datePanel *tview.InputField

	songsPanel      *tview.Table
	songActionPanel *tview.Flex

	outputPanel *tview.List
	outputChan  chan OutputMessage

	focusPrimitives   []tview.Primitive
	currentFocusIndex int

	keyBindings KeyBindings
}

func NewTUI() *TUI {
	ui := &TUI{
		client:            api.New(),
		currentFocusIndex: -1,
		keyBindings:       NewKeyBinding(),
		outputChan:        make(chan OutputMessage, 100),
	}

	ui.app = tview.NewApplication()

	ui.yearPanel = ui.createYearPanel()
	ui.datePanel = ui.createDatePanel()
	ui.songsPanel = ui.createSongsPanel()
	ui.outputPanel = ui.createOutputPanel()

	ui.leftPanel = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ui.yearPanel, 0, 5, false).
		AddItem(ui.datePanel, 0, 5, false)

	ui.rightPanel = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(ui.songsPanel, 0, 7, false).
		AddItem(ui.outputPanel, 0, 3, false)

	ui.layout = tview.NewFlex()
	ui.layout.AddItem(ui.leftPanel, 0, 3, false)
	ui.layout.AddItem(ui.rightPanel, 0, 7, false)

	ui.focusPrimitives = append(ui.focusPrimitives, ui.yearPanel, ui.datePanel)

	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		name := ui.keyBindings.SearchKey(event.Key())
		switch name {
		case "switch_focus":
			nextFocusIndex := ui.currentFocusIndex + 1
			if nextFocusIndex > len(ui.focusPrimitives)-1 {
				nextFocusIndex = 0
			}

			ui.app.SetFocus(ui.focusPrimitives[nextFocusIndex])
			ui.currentFocusIndex = nextFocusIndex

			return nil
		}

		return event
	})

	return ui
}

func (ui *TUI) Start() error {
	go func() {
		for out := range ui.outputChan {
			ui.outputPanel.AddItem(fmt.Sprintf("%d | [%s] %s", ui.outputPanel.GetItemCount(), time.Now().Format(time.RFC3339), out.Message), "", 0, nil)
			ui.outputPanel.SetCurrentItem(-1)
		}
	}()
	return ui.app.SetRoot(ui.layout, true).Run()
}

func (ui *TUI) createYearPanel() *tview.InputField {
	yearPanel := tview.NewInputField().SetLabel(" Year ")

	yearPanel.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		var text = yearPanel.GetText()
		if len(text) == 0 {
			return
		}

		songs, err := ui.client.Year100(text)
		if err != nil {
			ui.outputChan <- OutputMessage{Message: fmt.Sprintf("err: %s", err.Error())}
			return
		}

		for hcol, header := range year100Header {
			ui.songsPanel.SetCell(0, hcol, tview.NewTableCell(header))
		}
		for row, song := range songs {
			ui.songsPanel.SetCell(row+1, 0, tview.NewTableCell(fmt.Sprintf(" %s", song.Rank)))
			ui.songsPanel.SetCell(row+1, 1, tview.NewTableCell(song.Song))
			ui.songsPanel.SetCell(row+1, 2, tview.NewTableCell(song.Artist))
		}

		ui.focusPrimitives = append(ui.focusPrimitives, ui.songsPanel)

		yearPanel.SetText("")
	})

	yearPanel.SetBorder(true).SetTitle(" Year Input ")
	return yearPanel
}

func (ui *TUI) createDatePanel() *tview.InputField {
	dataPanel := tview.NewInputField().SetLabel(" Date ")
	dataPanel.SetBorder(true).SetTitle(" Date Input ")
	return dataPanel
}

func (ui *TUI) createSongsPanel() *tview.Table {
	songsPanel := tview.NewTable()
	songsPanel.SetBorder(true).SetTitle(" songs ")
	return songsPanel
}

func (ui *TUI) createOutputPanel() *tview.List {
	outputPanel := tview.NewList()
	outputPanel.SetBorder(true).SetTitle(" output ")
	return outputPanel
}
