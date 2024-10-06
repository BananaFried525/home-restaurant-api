package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(c *gin.Context, recovered any) {
	if err, ok := recovered.(string); ok {
		log.Printf("error: %s", err)
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "50001", "message": "INTERNAL_SERVER_ERROR"})
}
