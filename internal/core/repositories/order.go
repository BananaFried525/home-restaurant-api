package repositories

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) ports.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) BulkCreate(data []entities.Order) (*[]entities.Order, error) {
	result := data
	err := o.db.Model(&entities.Order{}).Create(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
