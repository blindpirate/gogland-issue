package main

import (
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Run([]string{"ls"})
}
