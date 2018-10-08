package cmd

import (
	"fmt"

	"github.com/junland/warehouse/server"
	flag "github.com/spf13/pflag"
)

// BinVersion describes built binary version.
var BinVersion string

// GoVersion Describes Go version that was used to build the binary.
var GoVersion string

var (
	version, help bool
)

// init defines configuration flags and environment variables.
func init() {
	flags := flag.CommandLine
	flags.BoolVarP(&help, "help", "h", false, "Show this help")
	flags.BoolVar(&version, "version", false, "Display version information")
	flags.SortFlags = false
	flag.Parse()
}

// PrintHelp prints help text.
func PrintHelp() {
	fmt.Printf("Usage: warehouse [options] <command> [<args>]\n")
	fmt.Printf("\n")
	fmt.Printf(" Binary distribution service for people .\n")
	fmt.Printf("\n")
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	fmt.Printf("\n")
}

// PrintVersion prints version information about the binary.
func PrintVersion() {
	fmt.Printf("Made with love.\n")
	fmt.Printf("Version: %s\n", BinVersion)
	fmt.Printf("Go Version %s\n", GoVersion)
	fmt.Printf("License: MIT\n")
}

// Run is the entry point for starting the command line interface.
func Run() {

	if version {
		PrintVersion()
		return
	}

	if help {
		PrintHelp()
		return
	}

	envconfig := GetEnvConf()

	server.Start(envconfig)
}
