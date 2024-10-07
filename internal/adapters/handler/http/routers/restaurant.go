package routers

import (
	adapter "github.com/BananaFried525/home-restaurant-api/internal/adapters/handler/http"
	"github.com/BananaFried525/home-restaurant-api/internal/adapters/repository"
	"github.com/BananaFried525/home-restaurant-api/internal/core/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Restaurant(g *gin.RouterGroup, db *gorm.DB) {
	orderRepo := repository.NewTableOrderRepository(db)
	tableRepo := repository.NewTableRepository(db)
	tableOrderRepo := repository.NewTableOrderRepository(db)

	r := g.Group("/restaurant")
	// Table
	tableService := services.NewTableService(tableRepo)
	tableAdapter := adapter.NewHttpTableHandler(tableService)
	r.POST("/table", tableAdapter.AddTable)
	r.GET("/table", tableAdapter.GetTable)
	r.GET("/table/detail", tableAdapter.GetTableDetail)

	//Order
	orderService := services.NewOrderService(tableOrderRepo, orderRepo)
	orderAdapter := adapter.NewHttpOrderHandler(orderService)
	r.POST("/order/table", orderAdapter.CreateTableOrder)
}
