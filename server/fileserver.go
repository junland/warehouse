package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type FileInfo struct {
	Name    string
	Dir     bool
	Size    int64
	LastMod string
}

type Listing struct {
	Directory string
	Items     []FileInfo
}

func (c *Config) fileServerHandler(w http.ResponseWriter, r *http.Request) {
	dir, err := c.SetBaseDir(r.URL.Path)
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	base, err := filepath.Abs(dir)
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ps := httprouter.ParamsFromContext(r.Context())

	// `req` variable stores absolute path of resource requested.
	req := base + ps.ByName("filepath")

	file, err := IsFile(req)
	if err != nil {
		switch {
		case os.IsPermission(err):
			log.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		case os.IsExist(err):
			log.Errorln(err)
			w.WriteHeader(http.StatusGone)
			return
		default:
			log.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if file {
		SetHeaderForFile(w, req)
		http.ServeFile(w, r, req)
		return
	}

	list, err := DirList(req)
	if err != nil {
		log.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = ParseAndExecuteTmpl(w, c.Template, deftmpl, list)
	if err != nil {
		log.Errorln("eeeee: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// SetBaseDir sets the base directory based on the first parameter of a url.
func (c *Config) SetBaseDir(url string) (string, error) {
	log.Debugf("Setting base dir for %s", url)
	switch urlslice := strings.Split(url, "/"); urlslice[1] {
	case "assets":
		log.Debugf("Chose: %s", urlslice[1])
		return c.AssetsDir, nil
	default:
		log.Debugf("Chose default.")
		return "", fmt.Errorf("could not match %s route to fileserver", urlslice)
	}
}
