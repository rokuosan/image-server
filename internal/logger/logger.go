package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
}

var logger *Logger

func init() {
	if logger == nil {
		logger = &Logger{
			slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		}
	}
}

func New() Logger {
	return *logger
}

func (l Logger) Debugf(format string, v ...interface{}) {
	l.Debug(fmt.Sprintf(format, v...))
}
