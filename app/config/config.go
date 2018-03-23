package config

import (
	"os"
)

// ConfigureEnv sets default environment variables if ENVIRONMENT is not set.
func ConfigureEnv() error {
	//Set all the values as environment variables
	if os.Getenv("ENVIRONMENT") == "" {
		os.Setenv("PORT", "8080")
		os.Setenv("ENVIRONMENT", "dev")
	}
	return nil
}

// CheckEnv checks if the required environment variables are set.
func CheckEnv() bool {
	if os.Getenv("ENVIRONMENT") == "" {
		return false
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}

	return true
}
