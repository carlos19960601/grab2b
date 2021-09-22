package command

import (
	"github.com/urfave/cli/v2"
	"github.com/zengqiang96/grab2b/internal/show2b"
)

var run = &cli.Command{
	Name:   "run",
	Action: show2b.RunAction,
}
