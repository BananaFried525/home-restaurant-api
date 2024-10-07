package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/BananaFried525/home-restaurant-api/internal/core/domain/models"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/BananaFried525/home-restaurant-api/internal/core/utils"
	"gorm.io/gorm"
)

type TableOrderRepository struct {
	db *gorm.DB
}

func NewTableOrderRepository(db *gorm.DB) ports.TableOrderRepository {
	return &TableOrderRepository{
		db: db,
	}
}

func (t *TableOrderRepository) CreateTableOrder(tableOrder models.TableOrder) (*models.TableOrder, error) {
	// start transaction
	var err error
	txn := t.db.Begin(&sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if txn.Error != nil {
		return nil, txn.Error
	}
	defer func() {
		log.Println("defer")
		if r := recover(); r != nil {
			txn.Rollback()
		} else if r := txn.Error; r != nil {
			txn.Rollback()
		} else if err != nil {
			txn.Rollback()
		} else {
			txn.Commit()
		}
	}()

	result := tableOrder
	if err = txn.Model(&models.TableOrder{}).Create(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *TableOrderRepository) GetLatestTableOrder(tableID uint) (*models.TableOrder, error) {
	result := models.TableOrder{
		TableInfoID: tableID,
		Status:      models.TableOrderStatusCheckedOut,
	}

	if err := t.db.Model(&models.TableOrder{}).Preload("Table").Last(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *TableOrderRepository) CountTableOrder() (int64, error) {
	var result int64
	startMonth, endMonth := utils.GetStartEndOfMonth()
	if err := t.db.Model(&models.TableOrder{}).Where(
		"created_at between ? and ?",
		startMonth.Format(time.RFC3339),
		endMonth.Format(time.RFC3339),
	).Count(&result).Error; err != nil {
		return 0, err
	}

	return result, nil
}
