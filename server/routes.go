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

	chain := alice.New(c.FileServe, CORS, Recovery)

	// Set the routes for the application.
	router.Handler("GET", "/assets/*filepath", chain.ThenFunc(c.fileBrowser))

	return router
}
