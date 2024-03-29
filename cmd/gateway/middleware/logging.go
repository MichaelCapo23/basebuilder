package middleware

import (
	"fmt"
	"time"

	"github.com/MichaelCapo23/basebuilder/pkg/project/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware(internalLogger *logging.InternalLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		ctx := logging.WithLogger(c.Request.Context(), internalLogger)
		c.Request = c.Request.Clone(ctx)

		c.Next()

		internalLogger.Logger.Desugar().Info(
			fmt.Sprintf("%s\t%s =>\t%d", c.Request.Method, c.Request.URL.String(), c.Writer.Status()),
			zap.Namespace("httpRequest"),
			zap.String("requestMethod", c.Request.Method),
			zap.String("protocol", c.Request.Proto),
			zap.String("requestUrl", c.Request.URL.String()),
			zap.Int("status", c.Writer.Status()),
			zap.String("duration", fmt.Sprintf("%fs", time.Since(start).Seconds())),
			zap.String("remoteIp", c.GetHeader("X-Forwarded-For")),
		)
	}
}
