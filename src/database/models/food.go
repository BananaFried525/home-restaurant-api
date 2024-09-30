package models

import (
	"gorm.io/gorm"
)

type FoodStatus string

const (
	FoodActive          FoodStatus = "active"
	FoodTemporaryRunout FoodStatus = "temporary_runout"
	FoodInactive        FoodStatus = "inactive"
)

type Food struct {
	gorm.Model
	Name              string
	DisplayImage      string
	Description       string `gorm:"type:text"`
	Price             int    `gorm:"type:DECIMAL(10,2);default:0"`
	Status            FoodStatus
	IsShow            bool `gorm:"default:false"`
	DiscountFixed     int  `gorm:"default:0"`
	DiscountPercent   int  `gorm:"default:0"`
	VatPercent        int  `gorm:"default:7"`
	ServiceFee        int  `gorm:"default:0"`
	ServiceFeePercent int  `gorm:"default:10"`
}

func (Food) TableName() string {
	return "food"
}

func GetFood(tx *gorm.DB, id *int) []Food {
	var food []Food

	if id != nil {
		tx.Where("id = ?", id)
	}

	tx.Find(&food)

	return food
}
