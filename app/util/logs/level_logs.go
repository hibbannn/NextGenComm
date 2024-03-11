package logs

import (
	"errors"
	"github.com/rs/zerolog"
	_ "github.com/rs/zerolog"
	_ "github.com/rs/zerolog/log"
)

// SetLevel sets the logging level based on a string value.
func (l *Logs) SetLevel(levelString string) error {
	if levelString == "" {
		return errors.New("log level cannot be empty")
	}
	level, err := zerolog.ParseLevel(levelString)
	if err != nil {
		return err // Return error to be handled by the caller
	}

	zerolog.SetGlobalLevel(level)

	return nil
}
