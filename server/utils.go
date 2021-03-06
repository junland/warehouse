package server

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// Pidfile is a struct that describes a PID file.
type Pidfile struct {
	Name string
}

// CreatePID creates a new PID file.
func CreatePID(name string) *Pidfile {
	log.Debug("Creating and opening PID file...")

	if _, err := os.Stat(name); !os.IsNotExist(err) {
		// file exists
		value, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatalf("pidfile: failed to read pid ", err)
		}

		pid, err := strconv.Atoi(string(value))
		if err != nil {
			log.Fatalf("pidfile: failed to convert string to int ", err)
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			log.Info("Existing PID file does not have a running process, attempting to remove.")
			err := os.Remove(name)
			if err != nil {
				log.Error("pidfile: could not remove existing pidfile ", err)
				os.Exit(1)
			}
			log.Info("Removal complete...")
		} else {
			if err := process.Signal(syscall.Signal(0)); err == nil {
				log.Fatalf("Process %d is already running.", pid)
			}
		}
	}

	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("pidfile: failed to open pid ", err)
		os.Exit(1)
	}

	defer file.Close()

	log.Debug("Writing PID to PID file...")

	_, err = fmt.Fprintf(file, "%d", os.Getpid())
	if err != nil {
		log.Error("pidfile: failed to write pid to file ", err)
	}

	err = file.Close()
	if err != nil {
		log.Error("pidfile: failed to close pid file after writing to it ", err)
	}

	log.Debug("PID creation has been completed...")

	return &Pidfile{name}
}

// RemovePID removes the PID file.
func (pf *Pidfile) RemovePID() {
	log.Debug("Removing PID file...")

	err := os.Remove(pf.Name)
	if err != nil {
		log.Error("pidfile: failed to remove ", err)
	}
	log.Debug("PID file removed...")
}

// IsFile checks if a specified file is actually a file.
func IsFile(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.Mode().IsRegular(), err
}

// IsDir checks if a specified file is actually a directory.
func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.Mode().IsDir(), err
}

// SetHeaderForFile sets the `Content-Type` header for a file extension for a file.
func SetHeaderForFile(w http.ResponseWriter, file string) {
	switch ext := filepath.Ext(file); ext {
	case ".rpm":
		w.Header().Set("Content-Type", "application/x-rpm")
	case ".deb":
		w.Header().Set("Content-Type", "application/x-deb")
	default:
		w.Header().Set("Content-Type", "application/octet-stream")
	}
}

// ParseAndExecuteTmpl parses and executes a template.
func ParseAndExecuteTmpl(wr io.Writer, file string, fallback string, data interface{}) error {
	t := template.New("tmpl")

	if file == "" {
		t, err := t.Parse(fallback)
		if err != nil {
			return err
		}

		err = t.Execute(wr, data)
		if err != nil {
			return err
		}

		return nil
	}

	file, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	fdat, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	t, err = t.Parse(string(fdat))
	if err != nil {
		return err
	}

	err = t.Execute(wr, data)
	if err != nil {
		return err
	}

	return nil
}

// ByteCountBinary converts bytes to a human readable format
func ByteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
