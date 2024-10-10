package repositories

import (
	"database/sql"
	"errors"
	"log"

	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/BananaFried525/home-restaurant-api/internal/core/utils"
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

	_table := entities.TableInfo{}
	if err = txn.Model(&entities.TableInfo{}).Where("number=?", table.Number).First(&_table).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if _table.ID != 0 {
		err = utils.NewCustomError(utils.DataExistError)
		return err
	}

	data := entities.TableInfo{
		Number: table.Number,
		Status: entities.TableInfoStatusAvailable,
	}

	if err := txn.Model(&entities.TableInfo{}).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (t *TableRepository) GetTable(limit int, offset int) (*[]entities.TableInfo, error) {
	var result []entities.TableInfo
	if err := t.db.Model(&entities.TableInfo{}).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *TableRepository) GetTableByID(ID uint) (*entities.TableInfo, error) {
	var result entities.TableInfo
	if err := t.db.Model(&entities.TableInfo{}).Where("id = ?", ID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (t *TableRepository) UpdateTable(ID uint, data entities.TableInfo) error {
	if err := t.db.Model(&entities.TableInfo{}).Where("id = ?", ID).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}

func (t *TableRepository) DeltetTable(ID uint) error {
	if err := t.db.Model(&entities.TableInfo{}).Delete(&entities.TableInfo{ID: ID}).Error; err != nil {
		return err
	}

	return nil
}
