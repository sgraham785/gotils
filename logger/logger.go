package logger

import (
	"github.com/go-kit/kit/log"
)

// NewLogger takes a logger as a dependency
// and returns type Logger.
func NewLogger(logger log.Logger) *Logger {
	return &Logger{logger}
}

// Logger defines a logger
type Logger struct {
	logger log.Logger
}
