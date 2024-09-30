package models

import (
	"gorm.io/gorm"
)

type ReserveTable struct {
	gorm.Model
	TableID        int `gorm:"not null"`
	CustomerNumber int
	Table          *RestaurantTable `gorm:"foreignKey:TableID"`
	Order          []Order          `gorm:"foreignKey:ReserveTableID"`
}

func (ReserveTable) TableName() string {
	return "reserve_table"
}
