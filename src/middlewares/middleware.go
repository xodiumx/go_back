package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func ZapLoggerMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Handling request
		c.Next()

		// Request data logging
		logger.Info(
			fmt.Sprintf("method=%s path=%s status=%d client_ip=%s latency=%v",
				c.Request.Method, c.Request.URL.Path, c.Writer.Status(), c.ClientIP(), time.Since(start),
			),
		)
	}
}
