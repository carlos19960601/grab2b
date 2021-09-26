package show2b

import (
	"github.com/urfave/cli/v2"
)

func RunAction(ctx *cli.Context) error {
	ui := NewTUI()
	return ui.Start()
}
