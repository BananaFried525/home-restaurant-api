package services

import (
	"errors"
	"time"

	"github.com/BananaFried525/home-restaurant-api/internal/core/domain"
	"github.com/BananaFried525/home-restaurant-api/internal/core/entities"
	"github.com/BananaFried525/home-restaurant-api/internal/core/ports"
	"github.com/BananaFried525/home-restaurant-api/internal/core/utils"
	"gorm.io/gorm"
)

type OrderService struct {
	tableOrderRepo    ports.TableOrderRepository
	orderRepo         ports.OrderRepository
	customerOrderRepo ports.CustomerOrderRepository
}

func NewOrderService(tableOrderRepo ports.TableOrderRepository, orderRepo ports.OrderRepository, customerOrderRepo ports.CustomerOrderRepository) ports.OrderService {
	return &OrderService{
		tableOrderRepo:    tableOrderRepo,
		orderRepo:         orderRepo,
		customerOrderRepo: customerOrderRepo,
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
	tableOrderData := entities.TableOrder{
		Number:        tableOrderNumber,
		ReceiptNumber: tableOrderNumber,
		TableInfoID:   tableID,
		Status:        entities.TableOrderStatusOpen,
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

func (o *OrderService) CreateOrder(data domain.CustomerOrder) (domain.CustomerOrder, error) {
	result := domain.CustomerOrder{}

	now := time.Now()
	customerOrderData := entities.CustomerOrder{
		TableInfoID:  data.TableInfoID,
		TableOrderID: data.TableOrderID,
		OrderNumber:  now.Format("200602011504"),
		OrderedAt:    now,
	}
	customerOrder, err := o.customerOrderRepo.Create(customerOrderData)
	if err != nil {
		return result, nil
	}

	ordersData := make([]entities.Order, 0)
	for _, orderData := range data.Orders {
		tmpData := entities.Order{
			TableOrderID:    orderData.TableOrderID,
			CustomerOrderID: customerOrder.ID,
			FoodID:          orderData.FoodID,
			Status:          entities.OrderPending,
			PendingAt:       time.Now(),
		}

		ordersData = append(ordersData, tmpData)
	}

	orders, err := o.orderRepo.BulkCreate(ordersData)
	if err != nil {
		return result, nil
	}

	ordersResult := make([]domain.Order, 0)
	for _, order := range *orders {
		orderResult := domain.Order{
			ID:              order.ID,
			TableOrderID:    order.TableOrderID,
			CustomerOrderID: order.CustomerOrderID,
			FoodID:          order.FoodID,
			Status:          string(order.Status),
			PendingAt:       order.PendingAt.Format(time.RFC3339),
		}
		ordersResult = append(ordersResult, orderResult)
	}

	result = domain.CustomerOrder{
		ID:           customerOrder.ID,
		TableInfoID:  customerOrder.TableInfoID,
		TableOrderID: customerOrder.TableOrderID,
		OrderNumber:  customerOrder.OrderNumber,
		OrderAt:      customerOrder.OrderedAt.Format(time.RFC3339),
		Orders:       ordersResult,
	}

	return result, nil
}

func (o *OrderService) ViewOrder(customerID uint) (domain.CustomerOrder, error) {
	result := domain.CustomerOrder{}

	// get customerOrderByID include order
	customerOrder, err := o.customerOrderRepo.GetDetailByID(customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return result, errors.New("NOT FOUND")
		}

		return result, err
	}

	if customerOrder == nil || customerOrder.Orders == nil {
		return result, errors.New("NOT FOUND")
	}

	//format and return
	orders := make([]domain.Order, 0)
	for _, tmpOrder := range *customerOrder.Orders {
		var doneAt string
		if tmpOrder.DoneAt != nil {
			doneAt = tmpOrder.DoneAt.Format(time.RFC3339)
		}
		var cancelAt string
		if tmpOrder.DoneAt != nil {
			cancelAt = tmpOrder.CancelAt.Format(time.RFC3339)
		}

		tmp := domain.Order{
			ID:              tmpOrder.ID,
			TableOrderID:    tmpOrder.TableOrderID,
			CustomerOrderID: tmpOrder.CustomerOrderID,
			FoodID:          tmpOrder.FoodID,
			Status:          string(tmpOrder.Status),
			PendingAt:       tmpOrder.PendingAt.Format(time.RFC3339),
			DoneAt:          &doneAt,
			CancelAt:        &cancelAt,
			Remark:          tmpOrder.Remark,
			Food: domain.Food{
				ID:           tmpOrder.Food.ID,
				Name:         tmpOrder.Food.Name,
				DisplayImage: tmpOrder.Food.DisplayImage,
				Description:  tmpOrder.Food.Description,
				Price:        tmpOrder.Food.Price,
			},
		}
		orders = append(orders, tmp)
	}

	result = domain.CustomerOrder{
		ID:           customerOrder.ID,
		TableInfoID:  customerOrder.TableInfoID,
		TableOrderID: customerOrder.TableOrderID,
		CustomerID:   customerOrder.CustomerID,
		OrderNumber:  customerOrder.OrderNumber,
		OrderAt:      customerOrder.OrderedAt.Format(time.RFC3339),
		Remark:       customerOrder.Remark,
		Orders:       orders,
	}

	return result, nil
}
