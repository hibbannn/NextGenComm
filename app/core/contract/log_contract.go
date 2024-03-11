package contract

type LogContract interface {
	SetLevel(levelString string) error
	Close() error
	Debug(message string)
	Info(message string)
	Infof(format string, v ...interface{})
	Warn(message string)
	Warnf(format string, v ...interface{})
	Error(message string)
	Errorf(format string, v ...interface{})
	ErrorIfError(err error)
	Panic(message string)
	Panicf(format string, v ...interface{})
	Fatal(message string)
	Fatalf(format string, v ...interface{})
}
