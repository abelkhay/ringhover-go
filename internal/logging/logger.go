package logging

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	l    *zap.Logger
	once sync.Once
)

func L() *zap.Logger {
	once.Do(func() {
		var err error
		l, err = zap.NewProduction()
		if err != nil { panic(err) }
	})
	return l
}

func ZapMiddleware() gin.HandlerFunc {
	log := L()
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info("http",
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
		)
	}
}
