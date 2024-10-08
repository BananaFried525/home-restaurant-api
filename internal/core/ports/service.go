package ports

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
)

type TableService interface {
	AddTable(Number int) error
	GetListTable(limit int, offset int) ([]domain.Table, error)
	GetTableDetail(ID uint) (*entities.TableInfo, error)
}

type OrderService interface {
	CreateTableOrder(tableID uint) (domain.TableOrder, error)
	ViewMenu() ([]domain.Food, error)
	CreateOrder(domain.CustomerOrder) (domain.CustomerOrder, error)
	ViewOrder(customerOrderID uint) (domain.CustomerOrder, error)
}
