package command

import (
	"github.com/urfave/cli/v2"
	"github.com/zengqiang96/grab2b/internal/show2b"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "show2b"
	app.Description = "billboard tui"

	app.Action = show2b.RunAction

	return app
}
