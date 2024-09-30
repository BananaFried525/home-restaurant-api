package routers

import "github.com/gin-gonic/gin"

func New(r *gin.Engine) {
	v1 := r.Group("/v1")

	Restaurant(v1)
}
