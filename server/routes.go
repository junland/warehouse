package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	log "github.com/sirupsen/logrus"
)

// RegisterRoutes sets all the configured routes for the server to the designated handler and middleware.
func (c *Config) RegisterRoutes() *httprouter.Router {
	log.Debug("Setting route info...")

	// Set the router.
	router := httprouter.New()

	router.HandleMethodNotAllowed = true
	router.HandleOPTIONS = true
	router.RedirectTrailingSlash = true

	chain := alice.New(CORS, Recovery)

	// Set the asset route, this will be the default fileserver.
	router.Handler("GET", "/assets/*filepath", chain.ThenFunc(c.fileServerHandler))

	// Set optional routes.
	if c.RPMsDir != "" {
		router.Handler("GET", "/rpm/*filepath", chain.ThenFunc(c.fileServerHandler))
	}

	return router
}
