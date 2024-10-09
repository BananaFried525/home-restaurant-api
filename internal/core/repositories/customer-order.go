package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"gorm.io/gorm"
)

type CustomerOrderRepository struct {
	db *gorm.DB
}

func (c *CustomerOrderRepository) Create(data entities.CustomerOrder) (*entities.CustomerOrder, error) {
	// start transaction
	var err error
	txn := c.db.Begin(&sql.TxOptions{Isolation: sql.LevelReadCommitted})
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

	result := entities.CustomerOrder{
		TableInfoID:  data.TableInfoID,
		TableOrderID: data.TableOrderID,
		CustomerID:   data.CustomerID,
		OrderNumber:  data.OrderNumber,
		OrderedAt:    time.Now(),
	}

	if err := txn.Model(&entities.CustomerOrder{}).Create(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *CustomerOrderRepository) GetByID(ID uint) (*entities.CustomerOrder, error) {
	var result entities.CustomerOrder
	err := c.db.Model(&entities.CustomerOrder{}).Where("id = ?", ID).First(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Update implements ports.CustomerOrderRepository.
func (c *CustomerOrderRepository) Update(ID uint, data entities.CustomerOrder) error {
	panic("unimplemented")
}

func NewCustomerOrderRepository(db *gorm.DB) ports.CustomerOrderRepository {
	return &CustomerOrderRepository{
		db: db,
	}
}

func (c *CustomerOrderRepository) GetDetailByID(ID uint) (*entities.CustomerOrder, error) {
	var result entities.CustomerOrder
	err := c.db.Model(&entities.CustomerOrder{}).Preload("Orders", func(_db *gorm.DB) *gorm.DB {
		return _db.Preload("Food")
	}).First(&result, ID).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
