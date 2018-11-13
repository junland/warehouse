package server

import "testing"

//	listing := Listing{RealPath: "../cmd", RelPath: "../cmd", Items: files}

func TestDirList(t *testing.T) {
	list, err := DirList("../cmd")
	if err != nil {
		t.Fatal(err)
	}

	files := []FileInfo{{Name: "main.go", Dir: false, Size: 3514, HumanSize: "3.5 KiB", LastMod: "2018-12-01 19:28"}, {Name: "utils.go", Dir: false, Size: 3514, HumanSize: "3.5 KiB", LastMod: "2018-12-01 19:28"}, {Name: "utils_test.go", Dir: false, Size: 3514, HumanSize: "3.5 KiB", LastMod: "2018-12-01 19:28"}}

	if list.Items[0].Name != files[0].Name {
		t.Fatal(err)
	}

	if list.Items[1].Name != files[1].Name {
		t.Fatal(err)
	}

	if list.Items[2].Name != files[2].Name {
		t.Fatal(err)
	}
}
