package middlewares

import (
	"log"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/BananaFried525/home-restaurant-api/internal/core/utils"
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

	utils.CustomErrorHandler(c, utils.NewCustomError(utils.InternalServerError))
}
