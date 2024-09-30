package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/BananaFried525/home-restaurant-api/src/database"
	"github.com/BananaFried525/home-restaurant-api/src/database/models"
	"github.com/BananaFried525/home-restaurant-api/src/validates"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var table = make(map[int]TableAttribute)

func GetTable(c *gin.Context) {
	data := validates.GetTable(c)
	log.Println(data)

	c.JSON(http.StatusOK, gin.H{
		"code":  "20000",
		"count": len(table),
		"data":  table,
	})
}

func AddTable(c *gin.Context) {
	data := validates.AddTable(c)

	if _, ok := table[data.TableNumber]; ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "40002", "message": "TABLE_ALREADY_EXISTS"})
		return
	}

	table[data.TableNumber] = TableAttribute{
		ID:             uuid.New().String(),
		TableNumber:    data.TableNumber,
		CustomerNumber: data.CustomerNumber,
		StartDate:      time.Now().Format(time.RFC3339),
		EndDate:        "",
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": table,
	})
}

func GetMenu(c *gin.Context) {
	data, err := validates.GetMenu(c)
	if err != nil {
		return
	}

	tableData, ok := table[data.TableNumber]
	if !ok || data.TableID != tableData.ID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": "40002", "message": "TABLE_NOT_FOUND"})
		return
	}

	menu := make([]MenuAttribute, 0)
	dbTxn := database.Begin()
	foods := models.GetFood(dbTxn, nil)
	for _, food := range foods {
		menu = append(menu, MenuAttribute{
			FoodID:          food.ID,
			FoodName:        food.Name,
			FoodPrice:       food.Price,
			FoodImage:       food.DisplayImage,
			FoodDescription: food.Description,
			FoodStatus:      string(food.Status),
		})
	}

	dbTxn.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": menu,
	})
}

func AddOrder(c *gin.Context) {
	_, err := validates.AddOrder(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
	})
}
