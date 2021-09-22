package command

import "github.com/urfave/cli/v2"

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "grab2b"
	app.Description = "抓取billboard榜单数据"

	app.Commands = []*cli.Command{
		hot100,
		artist100,
		year100,
	}

	return app
}
