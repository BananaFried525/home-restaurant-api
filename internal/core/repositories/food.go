package repositories

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"gorm.io/gorm"
)

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) ports.FoodRepository {
	return &FoodRepository{
		db: db,
	}
}

func (f *FoodRepository) Get() (*[]entities.Food, error) {
	var result []entities.Food
	if err := f.db.Model(&entities.Food{}).Where("is_show = ?", true).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (f *FoodRepository) GetByID(ID uint) (*entities.Food, error) {
	var result entities.Food
	err := f.db.Model(&entities.Food{}).First(&result, ID).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (f *FoodRepository) GetListByID(ID []uint) (*[]entities.Food, error) {
	var result []entities.Food
	err := f.db.Model(&entities.Food{}).Find(&result, ID).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
