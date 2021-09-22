package command

import (
	"github.com/urfave/cli/v2"
	"github.com/zengqiang96/grab2b/internal"
)

var year100 = &cli.Command{
	Name: "year100",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "year",
			Aliases: []string{"y"},
			Usage:   "年份，格式如 YYYY",
		},
	},
	Action: internal.Year100Action,
}
