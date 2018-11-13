package cmd

import (
	"log"
	"os"
)

// GetEnvString defines a environment variable with a specified name, fallback value.
// The return is a string value.
func GetEnvString(key, fallback string) string {
	if s := os.Getenv(key); s != "" {
		return s
	}
	return fallback
}

// GetEnvBool defines a environment variable with a specified name, fallback value.
// The return is either a true or false.
func GetEnvBool(key string, fallback bool) bool {
	switch os.Getenv(key) {
	case "true":
		return true
	case "false":
		return false
	default:
		return fallback
	}
}

// GetEnvStringReq checks if a key exists for a enviroment variable, if it's empty the program will panic and exit.
func GetEnvStringReq(key string) string {
	if s := os.Getenv(key); s != "" {
		log.Panicf("Please set %s before continuing.", key)
	}
	return os.Getenv(key)
}
