package ports

import (
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
)

type TableRepository interface {
	CreateTable(table domain.Table) error
	GetTable(limit int, offset int) (*[]entities.TableInfo, error)
	GetTableByID(ID uint) (*entities.TableInfo, error)
	UpdateTable(ID uint, table domain.Table) error
	DeltetTable(ID uint) error
}

type TableOrderRepository interface {
	CreateTableOrder(tableOrder entities.TableOrder) (*entities.TableOrder, error)
	GetLatestTableOrder(tableID uint) (*entities.TableOrder, error)
	CountTableOrder() (int64, error)
}
type OrderRepository interface {
	BulkCreate(data []entities.Order) (*[]entities.Order, error)
}

type CustomerOrderRepository interface {
	Create(data entities.CustomerOrder) (*entities.CustomerOrder, error)
	GetByID(ID uint) (*entities.CustomerOrder, error)
	Update(ID uint, data entities.CustomerOrder) error
	GetDetailByID(ID uint) (*entities.CustomerOrder, error)
}

type FoodRepository interface {
	Get() (*[]entities.Food, error)
	GetByID(ID uint) (*entities.Food, error)
	GetListByID(ID []uint) (*[]entities.Food, error)
}
