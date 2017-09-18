package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "RGE"
	app.Usage = "Solve RGE with Golang"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Renormalize!")
		return nil
	}

	app.Run(os.Args)
}
