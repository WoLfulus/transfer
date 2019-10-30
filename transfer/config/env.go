package config

import "os"

const (
	// EnvHtpasswdPath is an environment variable name
	EnvHtpasswdPath string = "TRANSFER_HTPASSWD_PATH"

	// EnvDebug is an environment variable name
	EnvDebug string = "DOCKER_TRANSFER_DEBUG"

	// EnvService is an environment variable name
	EnvService string = "TRANSFER_SERVICE"
)

// GetPasswordsFile returns the location of passwords file
func GetPasswordsFile() string {
	file := os.Getenv(EnvHtpasswdPath)
	if file == "" {
		return "/data/auth/htpasswd"
	}
	return file
}

// IsDebug checks if software is in debug mode
func IsDebug() bool {
	return os.Getenv(EnvDebug) != ""
}

// IsService tells if we're running inside docker service
func IsService() bool {
	return os.Getenv(EnvService) != ""
}

// IsClient tells if we're running as a CLI
func IsClient() bool {
	return !IsService()
}
