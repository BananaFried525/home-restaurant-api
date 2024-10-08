package http

import (
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HttpOrderControllers struct {
	orderService ports.OrderService
}

func NewHttpOrderControllers(orderService ports.OrderService) *HttpOrderControllers {
	return &HttpOrderControllers{orderService: orderService}
}

type CreateTableOrderRequest struct {
	TableID uint `json:"table_id" binding:"required"`
}

func (h *HttpOrderControllers) CreateTableOrder(c *gin.Context) {
	var req CreateTableOrderRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "BAD REQUEST"})
		return
	}

	result, err := h.orderService.CreateTableOrder(req.TableID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "INTERNAL SERVER ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
