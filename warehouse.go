// Package main is a simple wrapper of the real warehouse entrypoint package.
//
// This package should NOT be extended or modified in any way; to modify the
// warehouse binary, work in the `gitlab.com/junland/warehouse/cmd` package.
//
package main

import (
	warehouse "github.com/junland/warehouse/cmd"
)

func main() {
	warehouse.Run()
}
