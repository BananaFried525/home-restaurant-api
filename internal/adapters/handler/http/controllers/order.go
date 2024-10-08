package http

import (
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
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

type OrderAttribute struct {
	TableOrderID uint `json:"table_order_id" binding:"required"`
	FoodID       uint `json:"food_id" binding:"required"`
}
type CreateCustomerOrderRequest struct {
	TableInfoID  uint             `json:"table_id" binding:"required"`
	TableOrderID uint             `json:"table_order_id" binding:"required"`
	Orders       []OrderAttribute `json:"orders" binding:"required"`
}

func (h *HttpOrderControllers) CreateCustomerOrder(c *gin.Context) {
	var req CreateCustomerOrderRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "BAD REQUEST"})
		return
	}

	orders := make([]domain.Order, 0)
	for _, order := range req.Orders {
		tmpOrder := domain.Order{
			TableOrderID: order.TableOrderID,
			FoodID:       order.FoodID,
		}
		orders = append(orders, tmpOrder)
	}

	data := domain.CustomerOrder{
		TableInfoID:  req.TableInfoID,
		TableOrderID: req.TableOrderID,
		Order:        orders,
	}

	result, err := h.orderService.CreateOrder(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "INTERNAL SERVER ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
