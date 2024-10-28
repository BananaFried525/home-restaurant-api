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
	TableID       uint   `gorm:"not null"`
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
	Table          *Table           `gorm:"foreignKey:TableID"`
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
