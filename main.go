package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var errTooFewArgs = errors.New("too few arguments")

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
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 1 {
					return errTooFewArgs
				}
				fmt.Println("grammar file: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
