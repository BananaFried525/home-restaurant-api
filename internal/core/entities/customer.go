package entities

import (
	"time"

	"gorm.io/gorm"
)

const (
	CustomerStatusActive   CustomerStatus = "active"
	CustomerStatusInactive CustomerStatus = "inactive"
	CustomerStatusSuspend  CustomerStatus = "Suspend"
)

type CustomerStatus string

type Customer struct {
	ID           uint           `gorm:"primaryKey;autoIncrement:true"`
	Name         string         `gorm:"uniqle"`
	DisplyName   string         `gorm:"not null"`
	MobileNumber *string        `gorm:"size:10;index"`
	Status       CustomerStatus `gorm:"default:active"`
	CreatedAt    time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// association
	TableOrders    *[]TableOrder    `gorm:"foreignKey:CustomerID"`
	CustomerOrders *[]CustomerOrder `gorm:"foreignKey:CustomerID"`
}

func (Customer) TableName() string {
	return "customer"
}
