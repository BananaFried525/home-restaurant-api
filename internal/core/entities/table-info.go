package entities

import (
	"time"

	"gorm.io/gorm"
)

const (
	TableInfoStatusAvailable   TableInfoStatus = "available"
	TableInfoStatusUnavailable TableInfoStatus = "unavailable"
	TableInfoStatusReserved    TableInfoStatus = "reserved"
)

type TableInfoStatus string

type TableInfo struct {
	ID        uint            `gorm:"primaryKey;autoIncrement:true"`
	Number    int             `gorm:"not null;size:50;index"`
	Status    TableInfoStatus `gorm:"default:available"`
	CreatedAt time.Time       `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time       `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt  `gorm:"index"`

	// association
	TableOrders   *[]TableOrder    `gorm:"foreignKey:TableInfoID"`
	CustomerOrder *[]CustomerOrder `gorm:"foreignKey:TableInfoID"`
}

func (TableInfo) TableName() string {
	return "table_info"
}
