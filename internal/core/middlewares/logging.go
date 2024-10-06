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

func (l *LoggingMiddleware) SystemLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
