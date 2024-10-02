package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderPending OrderStatus = "pending"
	OrderDone    OrderStatus = "done"
	OrderCancel  OrderStatus = "cancel"
)

type Order struct {
	ID           uint        `gorm:"primaryKey;autoIncrement:true"`
	TableOrderID uint        `gorm:"not null"`
	FoodID       uint        `gorm:"not null"`
	Status       OrderStatus `gorm:"default:pending"`
	ConfirmDate  string
	FinishDate   string
	CancelDate   string
	CreatedAt    time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	TableOrder *TableOrder `gorm:"foreignKey:TableOrderID"`
	Food       *Food       `gorm:"foreignKey:FoodID"`
}

func (Order) TableName() string {
	return "order"
}

type BulkCreateOrderParams struct {
	Orders []Order
}

func BulkCreateOrder(params *BulkCreateOrderParams, dbTxn *gorm.DB) *[]Order {
	result := &params.Orders

	err := dbTxn.Create(&result).Error
	if err != nil {
		panic(err)
	}

	return result
}
