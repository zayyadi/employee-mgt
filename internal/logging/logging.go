package logging

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// InitLogger initializes and configures the logger
func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// Set formatter
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Set output
	logger.SetOutput(os.Stdout)

	// Set log level
	level, err := logrus.ParseLevel(strings.ToLower(os.Getenv("LOG_LEVEL")))
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	return logger
}
