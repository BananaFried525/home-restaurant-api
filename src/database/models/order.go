package models

import (
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderPending OrderStatus = "pending"
	OrderDone    OrderStatus = "done"
	OrderCancel  OrderStatus = "cancel"
)

type Order struct {
	gorm.Model
	ReserveTableID int         `gorm:"not null"`
	FoodID         int         `gorm:"not null"`
	Status         OrderStatus `gorm:"default:pending"`
	ConfirmDate    string
	FinishDate     string
	CancelDate     string
	ReserveTable   *ReserveTable `gorm:"foreignKey:ReserveTableID"`
	Food           *Food         `gorm:"foreignKey:FoodID"`
}

func (Order) TableName() string {
	return "order"
}
