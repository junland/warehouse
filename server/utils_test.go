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

func TestIsFile(t *testing.T) {
	dir, err := IsFile("./utils.go")
	if err != nil {
		t.Fatalf("Could not verify file: %s", err)
	}

	if !dir {
		t.Errorf("File is not a file for some reason")
	}
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

func TestByteCountBinary(t *testing.T) {
	size := ByteCountBinary(44)

	if size != "44 B" {
		t.Errorf("Wrong conversion: %s", size)
	}

	size = ByteCountBinary(4444)

	if size != "4.3 KiB" {
		t.Errorf("Wrong conversion: %s", size)
	}

	size = ByteCountBinary(4444444)

	if size != "4.2 MiB" {
		t.Errorf("Wrong conversion: %s", size)
	}

	size = ByteCountBinary(4444444444)

	if size != "4.1 GiB" {
		t.Errorf("Wrong conversion: %s", size)
	}

	size = ByteCountBinary(4444444444444)

	if size != "4.0 TiB" {
		t.Errorf("Wrong conversion: %s", size)
	}

	size = ByteCountBinary(4444444444444444)

	if size != "3.9 PiB" {
		t.Errorf("Wrong conversion: %s", size)
	}

	size = ByteCountBinary(4444444444444444444)

	if size != "3.9 EiB" {
		t.Errorf("Wrong conversion: %s", size)
	}
}
