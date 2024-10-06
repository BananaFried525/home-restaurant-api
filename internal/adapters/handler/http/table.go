package http

import (
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HttpTableHandler struct {
	service ports.TableService
}

func NewHttpTableHandler(service ports.TableService) *HttpTableHandler {
	return &HttpTableHandler{service: service}
}

type AddTableRequest struct {
	Number   int `json:"number" binding:"required"`
	Capacity int `json:"capacity" min:"1" binding:"required"`
}

func (h *HttpTableHandler) AddTable(c *gin.Context) {
	var req AddTableRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "BAD REQUEST"})
		return
	}

	if err := h.service.AddTable(req.Number); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "INTERNAL SERVER ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "SUCCESS"})
}
