package services

import (
	"time"

	"github.com/BananaFried525/home-restaurant-api/src/database/models"
	"gorm.io/gorm"
)

func GetTable(params GetTableInfoParams, dbTxn *gorm.DB) []models.TableInfo {
	var tableInfoList []models.TableInfo

	err := dbTxn.Limit(params.Limit).Offset(params.Offset).Find(&tableInfoList).Error
	if err != nil {
		panic(err)
	}

	return tableInfoList
}

func CountTableOrder(dbTxn *gorm.DB) (result int64) {
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	err := dbTxn.Model(&models.TableOrder{}).Where("created_at BETWEEN ? AND ?", startOfMonth.Format(time.RFC3339), endOfMonth.Format(time.RFC3339)).Count(&result).Error
	if err != nil {
		panic(err)
	}

	return
}
