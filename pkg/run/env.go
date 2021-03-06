package run

import (
	"os"
	"os/user"
)

const (
	// EnvInfrakitHome is the environment variable for defining the top level working directory
	// for infrakit.
	EnvInfrakitHome = "INFRAKIT_HOME"
)

// InfrakitHome returns the directory of INFRAKIT_HOME if specified. Otherwise, it will return
// the user's home directory.  If that cannot be determined, then it returns the current working
// directory.  If that still cannot be determined, a temporary directory is returned.
func InfrakitHome() string {
	dir := os.Getenv("INFRAKIT_HOME")
	if dir != "" {
		return dir
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	dir = os.Getenv("HOME")
	if dir != "" {
		return dir
	}
	dir, err := os.Getwd()
	if err == nil {
		return dir
	}
	return os.TempDir()
}

// GetEnv returns the value at the environment variable 'env'.  If the value is not found
// then default value is returned
func GetEnv(env string, defaultValue string) string {
	v := os.Getenv(env)
	if v != "" {
		return v
	}
	return defaultValue
}
