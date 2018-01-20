package command

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/urfave/cli"
)

var errTooFewArgs = errors.New("too few arguments")

// Serve runs a webserve that serves the contents of the web/dist directory
func Serve(c *cli.Context) error {
	if len(c.Args()) < 1 {
		return errTooFewArgs
	}
	// fileserver := http.StripPrefix("/dist/", http.FileServer(assetFS()))
	fileserver := injectDist(http.FileServer(assetFS()))
	fmt.Println("grammar file: ", c.Args().First())
	fmt.Println("http://localhost:8080/")
	err := http.ListenAndServe(":8080", fileserver)
	return err
}

func injectDist(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r2 := new(http.Request)
		*r2 = *r
		r2.URL = new(url.URL)
		*r2.URL = *r.URL
		r2.URL.Path = "dist/" + r2.URL.Path
		h.ServeHTTP(w, r2)
	})
}
