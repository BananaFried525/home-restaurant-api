package routers

import (
	adapter "github.com/BananaFried525/home-restaurant-api/internal/adapters/handler/http"
	"github.com/BananaFried525/home-restaurant-api/internal/adapters/repository"
	"github.com/BananaFried525/home-restaurant-api/internal/core/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Restaurant(g *gin.RouterGroup, db *gorm.DB) {
	r := g.Group("/restaurant")

	// Table
	tableRepo := repository.NewTableRepository(db)
	tableService := services.NewTableService(tableRepo)
	tableHandler := adapter.NewHttpTableHandler(tableService)
	r.POST("/table", tableHandler.AddTable)
}
