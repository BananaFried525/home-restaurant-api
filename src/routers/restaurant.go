package routers

import (
	"github.com/BananaFried525/home-restaurant-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func Restaurant(g *gin.RouterGroup) {
	r := g.Group("/restaurant")

	r.GET("/table", controllers.GetTableList)
	r.POST("/table-order", controllers.CreateTableOrder)
	// r.GET("/table-order", controllers.)
}
