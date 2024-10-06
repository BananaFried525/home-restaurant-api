package middlewares

import (
	"log"
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type RecoveryMiddleware struct {
}

func NewRecoverMiddleware() ports.RecoveryMiddleware {
	return &RecoveryMiddleware{}
}

func (r *RecoveryMiddleware) Recovery(c *gin.Context, recovered any) {
	if err, ok := recovered.(string); ok {
		log.Printf("error: %s", err)
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"code": "50001", "message": "INTERNAL_SERVER_ERROR"})
}
