package repository

import (
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
	if err := t.db.Model(&models.TableInfo{}).Create(&table).Error; err != nil {
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
