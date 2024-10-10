package http

import (
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/BananaFried525/home-restaurant-api/internal/core/utils"
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
		utils.CustomErrorHandler(c, utils.NewCustomError(utils.BadRequestError))
		return
	}

	result, err := h.orderService.CreateTableOrder(req.TableID)
	if err != nil {
		utils.CustomErrorHandler(c, err)
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
		utils.CustomErrorHandler(c, utils.NewCustomError(utils.BadRequestError))
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
		Orders:       orders,
	}

	result, err := h.orderService.CreateOrder(data)
	if err != nil {
		utils.CustomErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

type GetOrderDetailRequest struct {
	CustomerOrderID uint `form:"customer_order_id" binding:"required"`
}

func (h *HttpOrderControllers) GetOrderDetail(c *gin.Context) {
	var req GetOrderDetailRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.CustomErrorHandler(c, utils.NewCustomError(utils.BadRequestError))
		return
	}

	result, err := h.orderService.ViewOrder(req.CustomerOrderID)
	if err != nil {
		utils.CustomErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

type GetMenuRequest struct {
}

func (h *HttpOrderControllers) GetMenu(c *gin.Context) {
	var req GetMenuRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.CustomErrorHandler(c, utils.NewCustomError(utils.BadRequestError))
		return
	}

	result, err := h.orderService.ViewMenu()
	if err != nil {
		utils.CustomErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
