package entities

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableStatusAvailable   TableStatus = "available"
	TableStatusUnavailable TableStatus = "unavailable"
	TableStatusReserved    TableStatus = "reserved"
)

type TableStatus string

type Table struct {
	ID        uint           `gorm:"primaryKey;autoIncrement:true"`
	Number    int            `gorm:"not null;size:50;index"`
	Status    TableStatus    `gorm:"default:available"`
	CreatedAt time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// association
	TableOrders   *[]TableOrder    `gorm:"foreignKey:TableID"`
	CustomerOrder *[]CustomerOrder `gorm:"foreignKey:TableID"`
}

func (Table) TableName() string {
	return "table"
}
