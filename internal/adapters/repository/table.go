package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain/models"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"gorm.io/gorm"
)

type TableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) ports.TableRepository {
	return &TableRepository{db: db}
}

func (t *TableRepository) CreateTable(table domain.Table) error {
	// start transaction
	var err error
	txn := t.db.Begin(&sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if txn.Error != nil {
		return txn.Error
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

	_table := models.TableInfo{}
	if err = txn.Model(&models.TableInfo{}).Where("number=?", table.Number).First(&_table).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if _table.ID != 0 {
		err = errors.New("DATA EXIST")
		return err
	}

	data := models.TableInfo{
		Number: table.Number,
		Status: models.TableInfoStatusAvailable,
	}

	if err := txn.Model(&models.TableInfo{}).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (t *TableRepository) GetTable(limit int, offset int) (*[]models.TableInfo, error) {
	var result []models.TableInfo
	if err := t.db.Model(&models.TableInfo{}).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *TableRepository) GetTableByID(ID uint) (*models.TableInfo, error) {
	var result models.TableInfo
	if err := t.db.Model(&models.TableInfo{}).Where("id=?", ID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *TableRepository) UpdateTable(ID uint, table domain.Table) error {

	data := models.TableInfo{
		ID:     ID,
		Number: table.Number,
		Status: models.TableInfoStatus(table.Status),
	}

	if err := t.db.Model(&models.TableInfo{}).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (t *TableRepository) DeltetTable(ID uint) error {
	if err := t.db.Model(&models.TableInfo{}).Delete(&models.TableInfo{ID: ID}).Error; err != nil {
		return err
	}

	return nil
}
