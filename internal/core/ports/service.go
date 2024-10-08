package ports

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
)

type TableService interface {
	AddTable(Number int) error
	GetListTable(limit int, offset int) (*[]entities.TableInfo, error)
	GetTableDetail(ID uint) (*entities.TableInfo, error)
	UpdateTable(ID uint, table domain.Table) error
	DeleteTable(ID uint) error
}

type OrderService interface {
	CreateTableOrder(tableID uint) (domain.TableOrder, error)
}
