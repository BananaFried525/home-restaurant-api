package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerOrder struct {
	ID           uint `gorm:"primaryKey;autoIncrement:true"`
	TableInfoID  uint `gorm:"not null"`
	TableOrderID uint `gorm:"not null"`
	CustomerID   *uint
	OrderNumber  string         `gorm:"unique;size:12"`
	OrderedAt    time.Time      `gorm:"not null"`
	Remark       string         `gorm:"type:text"`
	CreatedAt    time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// association
	TableInfo  *TableInfo  `gorm:"foreignKey:TableInfoID"`
	TableOrder *TableOrder `gorm:"foreignKey:TableOrderID"`
	Orders     *[]Order    `gorm:"foreignKey:CustomerOrderID"`
	Customer   *Customer   `gorm:"foreignKey:CustomerID"`
}

func (CustomerOrder) TableName() string {
	return "customer_order"
}
