package models

import (
	"gorm.io/gorm"
)

const (
	TableStatusAvailable   TableStatus = "available"
	TableStatusUnavailable TableStatus = "unavailable"
)

type TableStatus string

type RestaurantTable struct {
	ID       uint        `gorm:"primaryKey;autoIncrement:true"`
	Name     string      `gorm:"not null"`
	Capacity int         `gorm:"not null"`
	Status   TableStatus `gorm:"default:available"`

	ReserveTables *[]ReserveTable `gorm:"foreignKey:TableID"`
}

func (RestaurantTable) TableName() string {
	return "restaurant_table"
}

type GetTableParams struct {
	Limit  int
	Offset int
}

func GetTable(param *GetTableParams, dbTxn *gorm.DB) ([]RestaurantTable, error) {
	var table []RestaurantTable
	r := dbTxn.Preload("ReserveTables").Limit(param.Limit).Offset(param.Offset).Find(&table)
	if r.Error != nil {
		return table, r.Error
	}

	return table, nil
}

type GetTableByIDParams struct {
	ID             uint
	ReserveTableID *uint
}

func GetTableByID(params *GetTableByIDParams, dbTxn *gorm.DB) *RestaurantTable {
	var table *RestaurantTable

	//prepare
	dbTxn = dbTxn.Preload("ReserveTables", func(_db1 *gorm.DB) *gorm.DB {
		_db1 = _db1.Where("id = ?", params.ReserveTableID)
		return _db1.Where("status in (?, ?)", ReserveTableStatusReserve, ReserveTableStatusDinning).Order("reserve_table.id desc").Limit(1)
	}).Where("restaurant_table.id = ?", params.ID)

	// exec
	err := dbTxn.First(&table).Error
	if err != nil {
		panic(err.Error())
	}

	return table
}

func UpdateTableStatus(tableID uint, status TableStatus, dbTxn *gorm.DB) error {
	r := dbTxn.Model(&RestaurantTable{}).Where("id = ?", tableID).Update("status", status)
	return r.Error
}
