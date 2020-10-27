package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func NewLogger() Logger {
	return Logger{Entry: logrus.NewEntry(logrus.New())}
}

const loggerKey = "logger"

// FromContext acquire logger from context
// If logger doesn't exist initialize default logger
func FromContext(ctx context.Context) Logger {
	val := ctx.Value(loggerKey)
	if val != nil {
		return Logger{Entry: val.(*logrus.Entry)}
	}

	logger := NewLogger()
	logger.Warn("context logger hasn't initialized on obtaining. Create default logger")
	return logger
}

//ToContext set logger to context
func ToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger.WithContext(ctx))
}
