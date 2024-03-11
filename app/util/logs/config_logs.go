package logs

import (
	"github.com/hibbannn/grpc-http-boilerplate/app/core/domain"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// Logs wraps the zerolog.Logger to provide a configurable logging setup.
// Logs struct definition remains the same as before
type Logs struct {
	config  domain.LoggingConfig
	logger  zerolog.Logger // Use direct logger instead of pointer
	logFile *os.File       // Track logFile for proper closure
}

// NewLogs creates a new Logs instance with the provided configuration.
func NewLogs(config domain.LoggingConfig) *Logs {
	return &Logs{config: config}
}

// InitializeLogs initializes the logger based on the provided configuration.
// Updated Init method
func (l *Logs) InitializeLogs() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	switch l.config.Format {
	case "console":
		output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		l.logger = zerolog.New(output).With().Timestamp().Logger()
	case "json":
		l.logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	case "file":
		var err error
		l.logFile, err = os.OpenFile(l.config.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		l.logger = zerolog.New(l.logFile).With().Timestamp().Logger()
	default:
		output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		l.logger = zerolog.New(output).With().Timestamp().Logger()
	}

	if err := l.SetLevel(l.config.Level); err != nil {
		return err
	}

	// Update global logger to the configured one
	log.Logger = l.logger

	return nil
}
