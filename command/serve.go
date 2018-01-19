package command

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/urfave/cli"
)

var errTooFewArgs = errors.New("too few arguments")

// Serve runs a webserve that serves the contents of the web/dist directory
func Serve(c *cli.Context) error {
	if len(c.Args()) < 1 {
		return errTooFewArgs
	}

	fmt.Println("grammar file: ", c.Args().First())
	fileserver := http.FileServer(http.Dir("web/dist"))
	fmt.Println("http://localhost:8080/")
	err := http.ListenAndServe(":8080", fileserver)
	return err
}
