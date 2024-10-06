package ports

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain/models"
)

type TableRepository interface {
	CreateTable(table domain.Table) error
	GetTable(limit int, offset int) (*[]models.TableInfo, error)
	// GetTableByID(ID uint) error
	// UpdateTable(ID uint, table domain.Table) error
	// DeltetTable(ID uint) error
}
