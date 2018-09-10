package logger

import (
	"context"
	"time"

	"github.com/twitchtv/twirp"

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

// TwirpLog implemetents logging middleware
func (l *Logger) TwirpLog(ctx context.Context, message string, input, output interface{}, err error, took time.Duration) error {
	pkg, _ := twirp.PackageName(ctx)
	service, _ := twirp.ServiceName(ctx)
	method, _ := twirp.MethodName(ctx)
	_ = l.logger.Log(
		"pkg", pkg,
		"service", service,
		"method", method,
		"@message", message,
		"input", input,
		"output", output,
		"err", err,
		"took", took,
		"@timestamp", time.Now().UTC(),
	)
	return nil
}
