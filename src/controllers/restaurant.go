package controllers

import (
	"net/http"

	"github.com/BananaFried525/home-restaurant-api/src/database"
	"github.com/BananaFried525/home-restaurant-api/src/helpers"
	"github.com/BananaFried525/home-restaurant-api/src/services"
	"github.com/BananaFried525/home-restaurant-api/src/validates"
	"github.com/gin-gonic/gin"
)

func GetTableList(c *gin.Context) {
	data, err := validates.GetTableList(c)
	if err != nil {
		return
	}

	dbTxn := database.Begin()
	tableInfoList := services.GetTable(services.GetTableInfoParams{Limit: data.Limit, Offset: data.Offset}, dbTxn)

	dbTxn.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": tableInfoList,
	})
}

func CreateTableOrder(c *gin.Context) {
	_, err := validates.CreateTableOrder(c)
	if err != nil {
		return
	}

	dbTxn := database.Begin()

	count := services.CountTableOrder(dbTxn)
	tableOrderNumber := helpers.PaddingIntToString(int(count)+1, 4, "0")

	dbTxn.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code": "20000",
		"data": tableOrderNumber,
	})
}
