package middleware

import (
	"context"

	"github.com/MichaelCapo23/jwtserver/pkg/project/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type traceIDKey string

const traceKey traceIDKey = "trace-id"

func TraceIDFromContext(ctx context.Context) string {
	v := ctx.Value(traceKey)
	if v == nil {
		return ""
	}

	t, ok := v.(string)
	if !ok {
		return ""
	}
	return t
}

func withTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceKey, id)
}

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		ctx := withTraceID(c.Request.Context(), traceID)
		logger := logging.FromContext(ctx).With("trace_id", traceID)
		internalLogger := logging.NewLogger(false)
		internalLogger.Logger = logger
		ctx = logging.WithLogger(ctx, internalLogger)
		c.Request = c.Request.Clone(ctx)
		c.Next()
	}
}
