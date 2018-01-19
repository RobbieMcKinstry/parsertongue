package main

import (
	"fmt"
	"os"

	"github.com/RobbieMcKinstry/parsertongue/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "parse"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Parsertongue!")
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "parse serve <grammar.ebnf>",
			Action:  command.Serve,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
