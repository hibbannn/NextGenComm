package logs

import "github.com/rs/zerolog/log"

// Close cleans up resources used by Logs, particularly the log file.
func (l *Logs) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

func (l *Logs) Info(message string) {
	log.Info().Msg(message)
}

func (l *Logs) Infof(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func (l *Logs) Error(message string) {
	log.Error().Msg(message)
}

func (l *Logs) Errorf(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}

func (l *Logs) Warn(message string) {
	log.Warn().Msg(message)
}

func (l *Logs) Fatal(message string) {
	log.Fatal().Msg(message)
}

func (l *Logs) Fatalf(format string, v ...interface{}) {
	log.Fatal().Msgf(format, v...)
}

func (l *Logs) Panicf(format string, v ...interface{}) {
	log.Panic().Msgf(format, v...)
}

func (l *Logs) Panic(message string) {
	log.Panic().Msg(message)
}

func (l *Logs) Warnf(format string, v ...interface{}) {
	log.Warn().Msgf(format, v...)
}

func (l *Logs) Debug(message string) {
	log.Debug().Msg(message)
}

func (l *Logs) Debugf(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}

func (l *Logs) FatalIfError(err error) {
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (l *Logs) PanicIfError(err error) {
	if err != nil {
		log.Panic().Err(err).Msg("")
	}
}

func (l *Logs) InfoIfError(err error) {
	if err != nil {
		log.Info().Err(err).Msg("")
	}
}

func (l *Logs) ErrorIfError(err error) {
	if err != nil {
		log.Error().Err(err).Msg("")
	}
}

func (l *Logs) WarnIfError(err error) {
	if err != nil {
		log.Warn().Err(err).Msg("")
	}
}

func (l *Logs) DebugIfError(err error) {
	if err != nil {
		log.Debug().Err(err).Msg("")
	}
}

func (l *Logs) FatalIfErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Fatal().Err(err).Msgf(format, v...)
	}
}

func (l *Logs) PanicIfErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Panic().Err(err).Msgf(format, v...)
	}
}

func (l *Logs) InfoIfErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Info().Err(err).Msgf(format, v...)
	}
}

func (l *Logs) ErrorIfErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Error().Err(err).Msgf(format, v...)
	}
}

func (l *Logs) WarnIfErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Warn().Err(err).Msgf(format, v...)
	}
}

func (l *Logs) DebugIfErrorf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Debug().Err(err).Msgf(format, v...)
	}
}
