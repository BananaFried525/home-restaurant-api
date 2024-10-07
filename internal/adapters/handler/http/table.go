package http

import (
	"errors"
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		if err.Error() == "DATA EXIST" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "DATA EXIST"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "INTERNAL SERVER ERROR"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "SUCCESS"})
}

type GetTableRequest struct {
	Limit  int `form:"limit,default=10"`
	Offset int `form:"offset,default=0"`
}

func (h *HttpTableHandler) GetTable(c *gin.Context) {
	var req GetTableRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "BAD REQUEST"})
		return
	}

	tableList, err := h.service.GetListTable(req.Limit, req.Offset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "INTERNAL SERVER ERROR"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tableList,
	})
}

type GetTableDetailRequest struct {
	ID uint `form:"id" binding:"required"`
}

func (h *HttpTableHandler) GetTableDetail(c *gin.Context) {
	var req GetTableDetailRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "BAD REQUEST"})
		return
	}

	table, err := h.service.GetTableDetail(req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": "NOT FOUND"})
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "INTERNAL SERVER ERROR"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":     table.Number,
			"status": table.Status,
		},
	})
}
