package logging

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func New() (*zap.Logger, error) {
	return zap.NewProduction()
}

func ZapMiddleware() gin.HandlerFunc {
	logger, _ := New()
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info("http",
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}
