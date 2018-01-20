package parserhttp

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/RobbieMcKinstry/parsertongue/grammar"
	"github.com/urfave/cli"
)

// Handler returns the HTTP Handler for this router.
func Handler(c *cli.Context) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", fileserver())
	mux.Handle("/grammar", grammarHandler(c))
	return mux
}

func fileserver() http.Handler {
	return injectDist(http.FileServer(assetFS()))
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

func grammarHandler(c *cli.Context) http.Handler {
	filename, root := c.Args().Get(0), c.Args().Get(1)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonObj := grammar.NewJSONFile(filename, root)
		encoder := json.NewEncoder(w)
		encoder.Encode(jsonObj)
	})
}
