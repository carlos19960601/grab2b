package main

import (
	"fmt"
	"os"

	"github.com/zengqiang96/grab2b/cmd/show2b/command"
)

func main() {
	app := command.App()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "启动失败: %s\n", err)
		os.Exit(1)
	}
}
