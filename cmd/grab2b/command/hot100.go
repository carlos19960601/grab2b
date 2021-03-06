package command

import (
	"github.com/urfave/cli/v2"
	"github.com/zengqiang96/grab2b/internal"
)

var hot100 = &cli.Command{
	Name: "hot100",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "date",
			Aliases: []string{"d"},
			Usage:   "日期，格式如 YYYY-MM-DD",
		},
	},
	Action: internal.Hot100Action,
}
