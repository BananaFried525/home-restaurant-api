package entities

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableOrderStatusOpen       TableOrderStatus = "open"
	TableOrderStatusReserved   TableOrderStatus = "reserved"
	TableOrderStatusCancel     TableOrderStatus = "cancel"
	TableOrderStatusCheckedOut TableOrderStatus = "checked_out"
)

type TableOrderStatus string
type TableOrder struct {
	ID            uint   `gorm:"primaryKey;autoIncrement:true"`
	Number        string `gorm:"unique;size:10"`
	ReceiptNumber string `gorm:"unique;size:10"`
	TableInfoID   uint   `gorm:"not null"`
	CustomerID    *uint
	Status        TableOrderStatus `gorm:"default:open"`
	ReservedAt    *time.Time
	CancelAt      *time.Time
	OpenedAt      *time.Time
	CheckedOutAt  *time.Time
	CreatedAt     time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	// association
	TableInfo      *TableInfo       `gorm:"foreignKey:TableInfoID"`
	Customer       *Customer        `gorm:"foreignKey:CustomerID"`
	CustomerOrders *[]CustomerOrder `gorm:"foreignKey:TableOrderID"`
}

func (TableOrder) TableName() string {
	return "table_order"
}

type CreateOrderTableParams struct {
	TableID uint
	Number  string
}

func CreateOrderTable(params *CreateOrderTableParams, dbTxn *gorm.DB) *TableOrder {
	tableOrder := &TableOrder{
		TableInfoID: params.TableID,
		Number:      params.Number,
	}

	err := dbTxn.Create(&tableOrder).Error
	if err != nil {
		panic(err)
	}

	return tableOrder
}
