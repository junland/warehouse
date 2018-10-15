// Package main is a simple wrapper of the real sled entrypoint package.
//
// This package should NOT be extended or modified in any way; to modify the
// sled binary, work in the `gitlab.com/<USER>/warehouse/cmd` package.
//
package main

import (
	warehouse "github.com/junland/warehouse/cmd"
)

func main() {
	warehouse.Run()
}
