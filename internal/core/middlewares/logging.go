package middlewares

import (
	"log"
	"time"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type LoggingMiddleware struct {
}

func NewLoggingMiddleware() ports.LoggingMiddleware {
	return &LoggingMiddleware{}
}

func (l *LoggingMiddleware) SystemRequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		requestPath := c.Request.URL.Path
		requestMethod := c.Request.Method

		log.Printf("START| %s %s \n", requestMethod, requestPath)

		// before request
		c.Next()

		// after request
		// access the status we are sending
		latency := time.Since(t).String()
		status := c.Writer.Status()
		log.Printf("END| %s %s %d %s \n", requestMethod, requestPath, status, latency)
	}
}
