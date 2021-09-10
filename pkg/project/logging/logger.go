package logging

import (
	"context"

	"github.com/MichaelCapo23/jwtserver/pkg/project"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const loggerKey = contextKey("logger")

type InternalLogger struct {
	Logger *zap.SugaredLogger
}

func NewLogger(isDebug bool) *InternalLogger {
	var l *zap.Logger
	if isDebug {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		l, _ = config.Build()
	} else {
		l, _ = zap.NewProduction()
	}
	zap.ReplaceGlobals(l)
	logger := l.WithOptions(zap.AddCallerSkip(1)).Sugar()

	return &InternalLogger{
		Logger: logger,
	}
}

func (l *InternalLogger) ErrorCtx(ctx context.Context, msg string, keysAndValues ...interface{}) {
	traceID := project.TraceFromCtx(ctx)
	if traceID == "" {
		l.Logger.Errorw(msg, keysAndValues...)
	} else {
		keysAndValues = append(keysAndValues, "trace_id")
		keysAndValues = append(keysAndValues, traceID)
		l.Logger.Errorw(msg, keysAndValues...)
	}
}

func (l *InternalLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.Logger.Infow(msg, keysAndValues...)
}

func (l *InternalLogger) Sync() error {
	return l.Logger.Sync()
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerKey).(*zap.SugaredLogger); ok {
		return logger
	}
	return nil
}

func WithLogger(ctx context.Context, logger *InternalLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger.Logger)
}
