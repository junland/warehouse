package server

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestServerShutdown(t *testing.T) {
	config := Config{
		LogLvl: "DEBUG",
		Port:   "0",
		PID:    "./test-server.pid",
		TLS:    false,
		Cert:   "",
		Key:    "",
	}

	go func() {
		err := Start(config)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	time.Sleep(2 * time.Second)

	stop <- os.Interrupt
}
