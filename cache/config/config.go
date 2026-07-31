package config

import (
	"os"
	"time"
)

func LoadConfig() *Config {
	config := DefaultConfig()

	if port := os.Getenv("PORT"); port != "" {
		config.Port = port
	}

	if file := os.Getenv("DATA_FILE"); file != "" {
		config.DataFile = file
	}

	if cleanup := os.Getenv("CLEANUP_INTERVAL"); cleanup != "" {
		if d, err := time.ParseDuration(cleanup); err == nil {
			config.CleanupInterval = d
		}
	}

	if autosave := os.Getenv("AUTOSAVE_INTERVAL"); autosave != "" {
		if d, err := time.ParseDuration(autosave); err == nil {
			config.AutoSaveInterval = d
		}
	}

	return &config
}

// Config holds application configuration values.
type Config struct {
	Port             string
	DataFile         string
	CleanupInterval  time.Duration
	AutoSaveInterval time.Duration
}

// DefaultConfig returns a Config populated with default values.
func DefaultConfig() Config {
	return Config{
		Port:             "8080",
		DataFile:         "data.json",
		CleanupInterval:  time.Hour,
		AutoSaveInterval: 5 * time.Minute,
	}
}
