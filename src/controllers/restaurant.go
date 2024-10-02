package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/BananaFried525/home-restaurant-api/src/database"
	"github.com/BananaFried525/home-restaurant-api/src/database/models"
	"github.com/BananaFried525/home-restaurant-api/src/validates"
	"github.com/gin-gonic/gin"
)

var table = make(map[int]TableAttribute)

func GetTable(c *gin.Context) {
	data, err := validates.GetTable(c)
	if err != nil {
		return
	}

	dbTxn := database.Begin()
	tables, err := models.GetTable(
		&models.GetTableParams{
			Limit:  data.Limit,
			Offset: data.Offset,
		},
		dbTxn,
	)
	if err != nil {
		log.Panic(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "50001", "message": "INTERNAL_SERVER_ERROR"})
		return
	}
	dbTxn.Commit()

	tablesFormat := make([]map[string]interface{}, 0)
	for _, table := range tables {
		_table := gin.H{
			"id":                table.ID,
			"table_name":        table.Name,
			"customer_capacity": table.Capacity,
			"table_status":      string(table.Status),
		}
		tablesFormat = append(tablesFormat, _table)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  "20000",
		"count": len(tablesFormat),
		"data":  tablesFormat,
	})
}

func AddTable(c *gin.Context) {
	data, err := validates.AddTable(c)
	if err != nil {
		return
	}

	dbTxn := database.Begin()
	table := models.GetTableByID(&models.GetTableByIDParams{ID: data.TableNumber}, dbTxn)

	if table == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "40401", "message": "NOT_FOUND"})
		return
	}
	if table.Status == models.TableStatusUnavailable {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "40003", "message": "TABLE_UNAVAILABLE"})
		return
	}
	if table.Capacity < data.CustomerNumber {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "40004", "message": "CAPACITY_OVER"})
		return
	}

	nowTime := time.Now()
	reservedTable, err := models.CreateReserveTable(
		&models.CreateReserveTableParams{
			TableID:        data.TableNumber,
			CustomerNumber: data.CustomerNumber,
			ReserveAt:      nowTime,
			DinningAt:      nowTime,
		},
		dbTxn,
	)
	if err != nil {
		log.Panic(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "50002", "message": "INTERNAL_SERVER_ERROR"})
		return
	}
	models.UpdateTableStatus(data.TableNumber, models.TableStatusUnavailable, dbTxn)

	dbTxn.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": reservedTable,
	})
}

func GetMenu(c *gin.Context) {
	_, err := validates.GetMenu(c)
	if err != nil {
		return
	}

	menu := make([]map[string]interface{}, 0)
	dbTxn := database.Begin()
	foods, err := models.GetFood(&models.GetFoodParams{}, dbTxn)
	if err != nil {
		log.Panic(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "50001", "message": "INTERNAL_SERVER_ERROR"})
		return
	}
	for _, food := range foods {

		menu = append(menu, gin.H{
			"id":          food.ID,
			"name":        food.Name,
			"price":       int(food.Price),
			"image":       food.DisplayImage,
			"description": food.Description,
			"status":      string(food.Status),
		})
	}

	dbTxn.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": menu,
	})
}

func GetTableDetail(c *gin.Context) {
	data, err := validates.GetTableDetail(c)
	if err != nil {
		return
	}

	dbTxn := database.Begin()
	table := models.GetTableByID(&models.GetTableByIDParams{ID: data.TableNumber}, dbTxn)
	if table == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "40401", "message": "NOT_FOUND"})
		return
	}
	dbTxn.Commit()

	var dinningDetail map[string]interface{}
	if table.Status == models.TableStatusUnavailable && len(*table.ReserveTables) != 0 {
		currentDinning := (*table.ReserveTables)[0]
		dinningDetail = gin.H{
			"reserve_table_id": currentDinning.ID,
			"customer_number":  currentDinning.CustomerNumber,
			"reserve_at":       currentDinning.ReserveAt.Format("2006-01-02 15:04:05"),
			"dinning_at":       currentDinning.DinningAt.Format("2006-01-02 15:04:05"),
			"status":           string(currentDinning.Status),
			"remark":           currentDinning.Remark,
		}
	}

	result := gin.H{
		"table_id":          table.ID,
		"table_name":        table.Name,
		"customer_capacity": table.Capacity,
		"table_status":      string(table.Status),
		"current_dinning":   dinningDetail,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": result,
	})
}

func OrderFood(c *gin.Context) {
	data, err := validates.OrderFood(c)
	if err != nil {
		return
	}

	dbTxn := database.Begin()

	table := models.GetReserveTableByID(&models.GetReserveTableByIDParam{ID: data.ReserveTableID}, dbTxn)
	if table == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "40401", "message": "NOT_FOUND"})
		return
	}

	foodIDs := make([]uint, 0)
	for _, food := range data.Food {
		foodIDs = append(foodIDs, food.FoodID)
	}
	foodList := *models.GetListFoodByID(&models.GetListFoodByIDParams{ID: foodIDs}, dbTxn)
	if len(foodList) != len(data.Food) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"code": "40401", "message": "NOT_FOUND"})
	}

	tableOrder := models.CreateOrderTable(&models.CreateOrderTableParams{TableID: table.ID}, dbTxn)

	_orders := make([]models.Order, 0)
	for _, food := range foodList {
		order := models.Order{
			TableOrderID: tableOrder.ID,
			FoodID:       food.ID,
		}

		_orders = append(_orders, order)
	}
	orders := models.BulkCreateOrder(&models.BulkCreateOrderParams{Orders: _orders}, dbTxn)

	dbTxn.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": orders,
	})
}
