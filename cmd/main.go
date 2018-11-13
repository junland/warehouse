package cmd

import (
	"fmt"
	"log"

	"github.com/junland/warehouse/server"
	flag "github.com/spf13/pflag"
)

// BinVersion describes built binary version.
var BinVersion string

// GoVersion Describes Go version that was used to build the binary.
var GoVersion string

// Default parameters when program starts without flags or environment variables.
const (
	defLvl      = "debug"
	defAccess   = true
	defPort     = "8080"
	defPID      = "/var/run/WAREHOUSE.pid"
	defTLS      = false
	defCert     = ""
	defKey      = ""
	defAssetDir = "./"
	defRPMDir   = "./"
)

var (
	confLogLvl, confPort, confPID, confCert, confKey, confAssetsDir, confRPMDir, confDebDir, confTmpl string
	enableTLS, enableAccess, version, help                                                            bool
)

// init defines configuration flags and environment variables.
func init() {
	flags := flag.CommandLine

	// Server configuration
	flags.StringVar(&confLogLvl, "log-level", GetEnvString("WAREHOUSE_LOG_LEVEL", defLvl), "Specify log level for output.")
	flags.BoolVar(&enableAccess, "access-log", GetEnvBool("WAREHOUSE_ACCESS_LOG", defAccess), "Specify weather to run with or without HTTP access logs.")
	flags.StringVar(&confPort, "port", GetEnvString("WAREHOUSE_SERVER_PORT", defPort), "Starting server port.")
	flags.StringVar(&confPID, "pid-file", GetEnvString("WAREHOUSE_SERVER_PID", defPID), "Specify server PID file path.")
	flags.StringVar(&confTmpl, "tmpl-file", GetEnvString("WAREHOUSE_TMPL", ""), "Specify a template file for the global file browser.")

	// TLS configuration
	flags.BoolVar(&enableTLS, "tls", GetEnvBool("WAREHOUSE_TLS", defTLS), "Specify weather to run server in secure mode.")
	flags.StringVar(&confCert, "tls-cert", GetEnvString("WAREHOUSE_TLS_CERT", defCert), "Specify TLS certificate file path.")
	flags.StringVar(&confKey, "tls-key", GetEnvString("WAREHOUSE_TLS_KEY", defKey), "Specify TLS key file path.")

	// Dir configuration
	flags.StringVar(&confAssetsDir, "asset-dir", GetEnvString("WAREHOUSE_ASSET_DIR", defAssetDir), "Specify path for generic assets.")
	flags.StringVar(&confRPMDir, "rpm-dir", GetEnvString("WAREHOUSE_RPM_DIR", ""), "Specify path for rpm packages.")
	flags.StringVar(&confDebDir, "deb-dir", GetEnvString("WAREHOUSE_DEB_DIR", ""), "Specify path for deb packages.")

	flags.BoolVarP(&help, "help", "h", false, "Show this help")
	flags.BoolVar(&version, "version", false, "Display version information")
	flags.SortFlags = false
	flag.Parse()
}

// PrintHelp prints help text.
func PrintHelp() {
	fmt.Printf("Usage: warehouse [options] \n")
	fmt.Printf("\n")
	fmt.Printf("File and binary distribution service for people.\n")
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
	fmt.Printf("License: GPLv2\n")
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

	config := server.Config{
		LogLvl:    confLogLvl,
		Access:    enableAccess,
		Port:      confPort,
		PID:       confPID,
		TLS:       enableTLS,
		Cert:      confCert,
		Key:       confKey,
		AssetsDir: confAssetsDir,
		RPMDir:    confRPMDir,
		DebDir:    confDebDir,
		Template:  confTmpl,
	}

	err := server.Start(config)
	if err != nil {
		log.Fatalln(err)
	}
}
