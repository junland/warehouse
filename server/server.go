package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

// Config struct provides configuration fields for the server.
type Config struct {
	LogLvl    string
	Access    bool
	Port      string
	PID       string
	TLS       bool
	Cert      string
	Key       string
	AssetsDir string
	RPMDir    string
	DebDir    string
	Template  string
}

var stop = make(chan os.Signal)

// Start sets up and starts the main server application
func Start(c Config) error {
	// Get log level environment variable.
	envLvl, err := log.ParseLevel(c.LogLvl)
	if err != nil {
		log.Error("Invalid log level ", envLvl)
	} else {
		// Setup logging with Logrus.
		log.SetLevel(envLvl)
	}

	err = c.CheckServeDirs()
	if err != nil {
		log.Fatal(err)
	}

	if c.TLS {
		if c.Cert == "" || c.Key == "" {
			log.Fatal("Invalid TLS configuration, please pass a file path for both the key and certificate.")
		}
	}

	log.Info("Setting up server...")

	router := c.RegisterRoutes()

	log.Debug("Setting up logging...")

	srv := &http.Server{Addr: ":" + c.Port, Handler: AccessLogger(router, c.Access)}

	log.Debug("Starting server on port ", c.Port)

	go func() {
		if c.TLS {
			err := srv.ListenAndServeTLS(c.Cert, c.Key)
			if err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	log.Info("Serving on port " + c.Port + ", press CTRL + C to shutdown.")

	p := CreatePID(c.PID)

	signal.Notify(stop, os.Interrupt)

	log.Warn("After notify...")

	<-stop // wait for SIGINT

	log.Warn("Shutting down server...")

	p.RemovePID()

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second) // shut down gracefully, but wait no longer than 45 seconds before halting.

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to shutdown server %s", err)
	}

	return nil
}

// CheckServeDirs pre-checks user specified directories to make sure that they exist.
func (c *Config) CheckServeDirs() error {
	dirs := map[string]string{"assets": c.AssetsDir, "rpm": c.RPMDir, "deb": c.DebDir}

	log.Debugln("Checking dirs: ", dirs)
	for k, v := range dirs {
		if v == "" {
			log.Infof("Route for %s not configured...", k)
			continue
		}

		v, err := filepath.Abs(v)
		if err != nil {
			return fmt.Errorf("could not parse path of %s ", v)
		}
		_, err = IsDir(v)
		if err != nil {
			return fmt.Errorf("%s is not a directory for configuring the %s file server route, please check your configuration", v, k)
		}
	}

	return nil
}
