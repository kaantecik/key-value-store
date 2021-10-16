package config

import (
	"time"
)

const (
	// Name of the project.
	Name = "kv-store"

	// DefaultPort represents port of the server.
	DefaultPort = "5000"

	// DefaultHost represents host of the server.
	DefaultHost = "0.0.0.0"

	// DefaultSaveLocation is where project backup files are saved.
	DefaultSaveLocation = "/tmp/" + Name

	// DefaultSaveInterval represents interval when files are saved.
	DefaultSaveInterval = 10 * time.Minute

	// AllowedConfigExt represents extensions which app allowed.
	AllowedConfigExt = "json"

	// LogPath is where log files are saved.
	LogPath = DefaultSaveLocation + "/logs"

	// HTTPLogPath is where http logs are saved.
	HTTPLogPath = LogPath + "/http.log"

	// ErrorLogPath is where error logs are saved.
	ErrorLogPath = LogPath + "/error.log"

	// AppLogPath is where app logs are saved.
	AppLogPath = LogPath + "/app.log"
)
