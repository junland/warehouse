package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type FileInfo struct {
	Name      string
	Dir       bool
	Size      int64
	HumanSize string
	LastMod   string
}

type Listing struct {
	RealPath string
	RelPath  string
	Items    []FileInfo
}

type sortDirNameFirst Listing

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

	list.RelPath = r.URL.Path

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

func (l sortDirNameFirst) Len() int      { return len(l.Items) }
func (l sortDirNameFirst) Swap(i, j int) { l.Items[i], l.Items[j] = l.Items[j], l.Items[i] }

// Dont worry about cases.
func (l sortDirNameFirst) Less(i, j int) bool {
	return l.Items[i].Dir
}

// DirList lists the information about files inside a directory.
func DirList(file string) (Listing, error) {
	var list []FileInfo

	log.Debugf("Looking up dir: %s", file)
	f, err := os.Open(file)
	if err != nil {
		return Listing{}, err
	}
	fi, err := f.Stat()
	if err != nil {
		return Listing{}, err
	}
	defer f.Close()

	if fi.IsDir() {
		files, err := ioutil.ReadDir(file)
		if err != nil {
			return Listing{}, err
		}

		// Start going thru each file and do stuff.
		for _, f := range files {

			// file name
			name := f.Name()
			if f.IsDir() {
				name += "/"
			}

			// skip hidden files.
			if strings.HasPrefix(name, ".") {
				continue
			}

			// file type
			dir := f.IsDir()

			// file size
			size := f.Size()

			// human file size
			hsize := ByteCountBinary(size)

			// file last mod time
			mod := f.ModTime().Format("2006-01-02 15:04")

			list = append(list, FileInfo{Name: name, Dir: dir, Size: size, HumanSize: hsize, LastMod: mod})
		}
		fmt.Println("Before sort: ", list)
		sort.Sort(sortDirNameFirst(Listing{RealPath: file, Items: list}))
		fmt.Println("After sort: ", list)
		return Listing{RealPath: file, Items: list}, nil
	}

	return Listing{}, fmt.Errorf("%s is not a directory", fi)
}
