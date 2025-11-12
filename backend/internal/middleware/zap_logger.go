package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
		}
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				fields = append(fields, zap.Error(e.Err))
			}
		}
		switch {
		case statusCode >= 500:
			logger.Error("internal server error", fields...)
		case statusCode >= 400:
			logger.Warn("client error", fields...)
		default:
			logger.Info("incoming request", fields...)
		}
	}
}
