package models

import "gorm.io/gorm"

type RestaurantTable struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Capacity int    `gorm:"not null"`
}

func (RestaurantTable) TableName() string {
	return "restaurant_table"
}
