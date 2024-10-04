package services

import (
	"github.com/BananaFried525/home-restaurant-api/src/database/models"
	"gorm.io/gorm"
)

func CreateTableOrder(params CreateTableOrderParams, dbTxn *gorm.DB) models.TableOrder {
	tableOrder := models.TableOrder{
		Number:        params.Number,
		ReceiptNumber: params.ReceiptNumber,
		TableInfoID:   params.TableInfoID,
		CustomerID:    params.CustomerID,
		Status:        models.TableOrderStatus(params.Status),
		ReservedAt:    params.ReservedAt,
		CancelAt:      params.CancelAt,
		OpenedAt:      params.OpenedAt,
		CheckedOutAt:  params.CheckedOutAt,
	}

	err := dbTxn.Create(&tableOrder).Error
	if err != nil {
		panic(err)
	}

	return tableOrder
}
