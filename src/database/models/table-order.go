package models

import (
	"time"

	"gorm.io/gorm"
)

type TableOrder struct {
	ID             uint           `gorm:"primaryKey;autoIncrement:true"`
	ReserveTableID uint           `gorm:"not null"`
	OrderAt        time.Time      `gorm:"not null"`
	CreatedAt      time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	ReserveTable *ReserveTable `gorm:"foreignKey:ReserveTableID"`
	Orders       *[]Order      `gorm:"foreignKey:TableOrderID"`
}

func (TableOrder) TableName() string {
	return "table_order"
}

type CreateOrderTableParams struct {
	TableID uint
}

func CreateOrderTable(params *CreateOrderTableParams, dbTxn *gorm.DB) *TableOrder {
	tableOrder := &TableOrder{
		ReserveTableID: params.TableID,
		OrderAt:        time.Now(),
	}

	err := dbTxn.Create(&tableOrder).Error
	if err != nil {
		panic(err)
	}

	return tableOrder
}
