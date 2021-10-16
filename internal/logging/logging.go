package logging

import (
	"os"

	"github.com/kaantecik/key-value-store/internal/config"
	log "github.com/sirupsen/logrus"
)

var (
	HttpLogger  = SetLogger(config.HTTPLogPath)
	AppLogger   = SetLogger(config.AppLogPath)
	ErrorLogger = SetLogger(config.ErrorLogPath)
)

// SetLogger function gets a path and creates folder if file doesn't exist in that path.
// Then creates a logger object. Logger's output is set to this file.
func SetLogger(path string) *log.Logger {
	logger := log.New()

	logger.SetFormatter(&log.TextFormatter{ForceColors: true})
	logger.SetLevel(log.TraceLevel)

	if _, err := os.Stat(config.LogPath); os.IsNotExist(err) {
		os.Mkdir(config.LogPath, 0755)
	}

	f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	logger.SetOutput(f)
	return logger
}
