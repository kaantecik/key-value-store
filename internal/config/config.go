package config

import (
	"time"
)

const (
	Name = "kv-store"

	DefaultPort = "5000"

	DefaultHost = "0.0.0.0"

	AllowedConfigExt = "json"

	LogPath = "/tmp/" + Name + "/logs"

	HTTPLogPath = LogPath + "/http.log"

	ErrorLogPath = LogPath + "/error.log"

	AppLogPath = LogPath + "/app.log"

	DefaultSaveLocation = "/tmp/" + Name

	DefaultSaveInterval = 10 * time.Minute
)
