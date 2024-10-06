package ports

import "github.com/gin-gonic/gin"

type LoggingMiddleware interface {
	SystemLog() gin.HandlerFunc
}

type RecoveryMiddleware interface {
	Recovery(c *gin.Context, recovered any)
}
