package show2b

import (
	"github.com/rivo/tview"
	"github.com/urfave/cli/v2"
)

func RunAction(ctx *cli.Context) error {
	grid := tview.NewGrid().
		SetRows(0, 0).
		SetColumns(40, 0).
		SetBorders(true)

	grid.SetTitle(" show2b ")

	charts := tview.NewTextView().SetText("charts").SetTitleAlign(tview.AlignCenter)
	grid.AddItem(charts, 0, 0, 1, 1, 0, 0, false)

	yeteDropdown := tview.NewDropDown().SetLabel(" year/date: ").SetOptions([]string{"2019", "2020"}, nil)
	grid.AddItem(yeteDropdown, 1, 0, 1, 1, 0, 0, false)

	table := tview.NewTextView().SetText("table").SetTextAlign(tview.AlignCenter)

	grid.AddItem(table, 0, 1, 2, 1, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(yeteDropdown).Run(); err != nil {
		return err
	}
	return nil
}
