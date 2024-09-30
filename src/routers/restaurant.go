package routers

import (
	"github.com/BananaFried525/home-restaurant-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func Restaurant(g *gin.RouterGroup) {
	r := g.Group("/restaurant")
	r.GET("/table", controllers.GetTable)
	r.POST("/table", controllers.AddTable)
	r.GET("/menu", controllers.GetMenu)
	r.POST("/order", controllers.AddOrder)
}
