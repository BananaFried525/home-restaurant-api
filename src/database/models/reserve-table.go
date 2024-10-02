package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	ReserveTableStatusReserve  ReserveTableStatus = "reserve"
	ReserveTableStatusDinning  ReserveTableStatus = "dinning"
	ReserveTableStatusCheckout ReserveTableStatus = "checkout"
	ReserveTableStatusCancel   ReserveTableStatus = "cancel"
)

type ReserveTableStatus string

type ReserveTable struct {
	ID             uint `gorm:"primaryKey;autoIncrement:true"`
	TableID        uint `gorm:"not null"`
	CustomerNumber int
	Status         ReserveTableStatus `gorm:"default:reserve"`
	Remark         string             `gorm:"type:text"`
	ReserveAt      time.Time          `gorm:"default:null"`
	DinningAt      time.Time          `gorm:"default:null"`
	CreatedAt      time.Time          `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time          `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;onUpdate:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt     `gorm:"index"`

	RestaurantTable *RestaurantTable `gorm:"foreignKey:TableID"`
	TableOrder      *[]TableOrder    `gorm:"foreignKey:ReserveTableID"`
}

func (ReserveTable) TableName() string {
	return "reserve_table"
}

type CreateReserveTableParams struct {
	TableID        uint
	CustomerNumber int
	ReserveAt      time.Time
	DinningAt      time.Time
}

func CreateReserveTable(params *CreateReserveTableParams, dbTxn *gorm.DB) (*ReserveTable, error) {
	reserveTable := &ReserveTable{
		TableID:        params.TableID,
		CustomerNumber: params.CustomerNumber,
		Status:         ReserveTableStatusDinning,
		ReserveAt:      params.ReserveAt,
		DinningAt:      params.DinningAt,
	}
	r := dbTxn.Create(&reserveTable)
	if r.Error != nil {
		return nil, r.Error
	}

	return reserveTable, nil
}

type GetReserveTableByIDParam struct {
	ID uint
}

func GetReserveTableByID(param *GetReserveTableByIDParam, dbTxn *gorm.DB) *ReserveTable {
	var result *ReserveTable

	err := dbTxn.Where("id = ?", param.ID).Find(&result).Error
	if err != nil {
		panic(err)
	}

	return result
}
