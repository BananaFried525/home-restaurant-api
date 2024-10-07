package ports

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain/models"
)

type TableRepository interface {
	CreateTable(table domain.Table) error
	GetTable(limit int, offset int) (*[]models.TableInfo, error)
	GetTableByID(ID uint) (*models.TableInfo, error)
	UpdateTable(ID uint, table domain.Table) error
	DeltetTable(ID uint) error
}

type TableOrderRepository interface {
	CreateTableOrder(tableOrder models.TableOrder) (*models.TableOrder, error)
	GetLatestTableOrder(tableID uint) (*models.TableOrder, error)
	CountTableOrder() (int64, error)
}
type OrderRepository interface {
}
