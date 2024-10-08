package entities

import (
	"time"

	"gorm.io/gorm"
)

type FoodStatus string

const (
	FoodActive          FoodStatus = "active"
	FoodTemporaryRunout FoodStatus = "temporary_runout"
	FoodInactive        FoodStatus = "inactive"
)

type Food struct {
	ID                uint `gorm:"primaryKey;autoIncrement:true"`
	Name              string
	DisplayImage      string
	Description       string         `gorm:"type:text"`
	Price             float64        `gorm:"type:DECIMAL(10,2);default:0"`
	Status            FoodStatus     `gorm:"default:active"`
	IsShow            bool           `gorm:"default:false"`
	DiscountFixed     int            `gorm:"default:0"`
	DiscountPercent   int            `gorm:"default:0"`
	VatPercent        int            `gorm:"default:7"`
	ServiceFee        int            `gorm:"default:0"`
	ServiceFeePercent int            `gorm:"default:10"`
	CreatedAt         time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time      `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func (Food) TableName() string {
	return "food"
}

type GetFoodParams struct {
	// ID *uint
}

func GetFood(params *GetFoodParams, dbTxn *gorm.DB) ([]Food, error) {
	var food []Food

	err := dbTxn.Where("is_show = true").Find(&food).Error
	if err != nil {
		return food, err
	}

	return food, nil
}

type GetListFoodByIDParams struct {
	ID []uint
}

func GetListFoodByID(params *GetListFoodByIDParams, dbTxn *gorm.DB) *[]Food {
	var result *[]Food

	err := dbTxn.Where("id in ?", params.ID).Find(&result).Error
	if err != nil {
		panic(err.Error())
	}

	return result
}
