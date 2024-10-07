package http

import (
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HttpOrderHandler struct {
	orderService ports.OrderService
}

func NewHttpOrderHandler(orderService ports.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{orderService: orderService}
}

type CreateTableOrderRequest struct {
	TableID uint `json:"table_id" binding:"required"`
}

func (h *HttpOrderHandler) CreateTableOrder(c *gin.Context) {
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
