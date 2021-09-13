package middleware

import (
	"context"

	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type traceIDKey string

const traceKey = "TRACE_ID"

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

func TraceMiddleware(internalLogger *logging.InternalLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		ctx := withTraceID(c.Request.Context(), traceID)
		logger := logging.FromContext(ctx).With(traceKey, traceID)
		internalLogger.Logger = logger
		ctx = logging.WithLogger(ctx, internalLogger)
		c.Request = c.Request.Clone(ctx)
		c.Next()
	}
}
