package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (c *Config) fileBrowser(w http.ResponseWriter, r *http.Request) {
	ps := httprouter.ParamsFromContext(r.Context())

	// `path` variable has the beginning path stripped out.
	path := ps.ByName("filepath")
}
