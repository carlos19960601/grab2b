package command

import (
	"github.com/urfave/cli/v2"
	"github.com/zengqiang96/grab2b/internal"
)

var artist100 = &cli.Command{
	Name: "artist100",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "date",
			Aliases: []string{"d"},
			Usage:   "日期，格式如 YYYY-MM-DD",
		},
	},
	Action: internal.Artist100Action,
}
