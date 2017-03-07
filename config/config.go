package config

import (
	"os"
)

func GetPort() string {
	configuredPort := os.Getenv("PORT")
	if configuredPort == "" {
		return "3000"
	}

	return configuredPort
}
