package ports

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain/models"
)

type TableService interface {
	AddTable(Number int) error
	GetListTable(limit int, offset int) (*[]models.TableInfo, error)
	GetTableDetail(ID uint) (*models.TableInfo, error)
	UpdateTable(ID uint, table domain.Table) error
	DeleteTable(ID uint) error
}
