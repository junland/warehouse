package server

import (
	"os"
	"testing"
)

func TestCreateAndRemovePID(t *testing.T) {
	os.RemoveAll("./test-util.pid")

	pid := CreatePID("./test-util.pid")

	pid.RemovePID()

}

func TestIsDir(t *testing.T) {
	dir, err := IsDir("../server")
	if err != nil {
		t.Fatalf("Could not verify directory: %s", err)
	}

	if !dir {
		t.Errorf("File is not a directory for some reason")
	}
}

func TestIsFile(t *testing.T) {
	dir, err := IsFile("./utils.go")
	if err != nil {
		t.Fatalf("Could not verify file: %s", err)
	}

	if !dir {
		t.Errorf("File is not a file for some reason")
	}
}
