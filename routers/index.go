package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(r *gin.Engine, db *gorm.DB) {
	v1 := r.Group("/v1")

	Restaurant(v1, db)
}
