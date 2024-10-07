package services

import (
	"time"

	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/domain/models"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/BananaFried525/home-restaurant-api/internal/core/utils"
)

type OrderService struct {
	tableOrderRepo ports.TableOrderRepository
	orderRepo      ports.OrderRepository
}

func NewOrderService(tableOrderRepo ports.TableOrderRepository, orderRepo ports.OrderRepository) ports.OrderService {
	return &OrderService{
		tableOrderRepo: tableOrderRepo,
		orderRepo:      orderRepo,
	}
}

func (o *OrderService) CreateTableOrder(tableID uint) (domain.TableOrder, error) {
	result := domain.TableOrder{}

	// count total table order on table
	count, err := o.tableOrderRepo.CountTableOrder()
	if err != nil {
		return result, err
	}

	// create table order
	now := time.Now()
	tableOrderNumber := utils.CreateRunningNumber(int(count) + 1)
	tableOrderData := models.TableOrder{
		Number:        tableOrderNumber,
		ReceiptNumber: tableOrderNumber,
		TableInfoID:   tableID,
		Status:        models.TableOrderStatusOpen,
		OpenedAt:      &now,
	}
	tableOrder, err := o.tableOrderRepo.CreateTableOrder(tableOrderData)
	if err != nil {
		return result, err
	}

	openedAt := tableOrder.OpenedAt.Format(time.RFC3339)
	result = domain.TableOrder{
		ID:            tableOrder.ID,
		Number:        tableOrder.Number,
		ReceiptNumber: &tableOrder.ReceiptNumber,
		TableID:       tableOrder.TableInfoID,
		CustomerID:    tableOrder.CustomerID,
		Status:        string(tableOrder.Status),
		OpenedAt:      &openedAt,
	}

	return result, nil
}
