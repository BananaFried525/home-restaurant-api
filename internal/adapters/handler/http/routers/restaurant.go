package routers

import (
	controllers "github.com/BananaFried525/home-restaurant-api/internal/adapters/handler/http/controllers"
	"github.com/BananaFried525/home-restaurant-api/internal/core/repositories"
	"github.com/BananaFried525/home-restaurant-api/internal/core/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Restaurant(g *gin.RouterGroup, db *gorm.DB) {
	orderRepo := repositories.NewTableOrderRepository(db)
	tableRepo := repositories.NewTableRepository(db)
	tableOrderRepo := repositories.NewTableOrderRepository(db)

	r := g.Group("/restaurant")
	// Table
	tableService := services.NewTableService(tableRepo)
	tableControllers := controllers.NewHttpTableController(tableService)
	r.POST("/table", tableControllers.AddTable)
	r.GET("/table", tableControllers.GetTable)
	r.GET("/table/detail", tableControllers.GetTableDetail)

	//Order
	orderService := services.NewOrderService(tableOrderRepo, orderRepo)
	orderControllers := controllers.NewHttpOrderControllers(orderService)
	r.POST("/order/table", orderControllers.CreateTableOrder)
}
