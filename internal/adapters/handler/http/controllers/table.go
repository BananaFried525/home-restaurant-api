package http

import (
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/BananaFried525/home-restaurant-api/internal/core/utils"
	"github.com/gin-gonic/gin"
)

type HttpTableController struct {
	service ports.TableService
}

func NewHttpTableController(service ports.TableService) *HttpTableController {
	return &HttpTableController{service: service}
}

type AddTableRequest struct {
	Number   int `json:"number" binding:"required"`
	Capacity int `json:"capacity" min:"1" binding:"required"`
}

func (h *HttpTableController) AddTable(c *gin.Context) {
	var req AddTableRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		utils.CustomErrorHandler(c, utils.NewCustomError(utils.BadRequestError))
		return
	}

	if err := h.service.AddTable(req.Number); err != nil {
		utils.CustomErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "SUCCESS"})
}

type GetTableRequest struct {
	Limit  int `form:"limit,default=10"`
	Offset int `form:"offset,default=0"`
}

func (h *HttpTableController) GetTable(c *gin.Context) {
	var req GetTableRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.CustomErrorHandler(c, utils.NewCustomError(utils.BadRequestError))
		return
	}

	tableList, err := h.service.GetListTable(req.Limit, req.Offset)
	if err != nil {
		utils.CustomErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tableList,
	})
}

type GetTableDetailRequest struct {
	ID uint `form:"id" binding:"required"`
}

func (h *HttpTableController) GetTableDetail(c *gin.Context) {
	var req GetTableDetailRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.CustomErrorHandler(c, utils.NewCustomError(utils.BadRequestError))
		return
	}

	table, err := h.service.GetTableDetail(req.ID)
	if err != nil {
		utils.CustomErrorHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":     table.Number,
			"status": table.Status,
		},
	})
}
