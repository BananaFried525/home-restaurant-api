package domain

type Table struct {
	ID     uint   `json:"id"`
	Number int    `json:"number"`
	Status string `json:"status"`
}

type TableOrder struct {
	ID            uint    `json:"id"`
	Number        string  `json:"number"`
	ReceiptNumber *string `json:"receipt_number"`
	TableID       uint    `json:"table_id"`
	CustomerID    *uint   `json:"customer_id"`
	Status        string  `json:"stauts"`
	ReservedAt    *string `json:"reserverd_at"`
	CancelAt      *string `json:"cancel_at"`
	OpenedAt      *string `json:"open_at"`
	CheckedOutAt  *string `json:"check_out_at"`
}

type Food struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	DisplayImage string  `json:"display_name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Status       string  `json:"stutus"`
}

type Order struct {
	ID              uint    `json:"id"`
	TableOrderID    uint    `json:"table_order_id"`
	CustomerOrderID uint    `json:"customer_order_id"`
	FoodID          uint    `json:"food_id"`
	Status          string  `json:"status"`
	PendingAt       string  `json:"pending_at"`
	DoneAt          *string `json:"done_at"`
	CancelAt        *string `json:"cancel_at"`
	Remark          *string `json:"remark"`
	Food            Food    `json:"food"`
}

type CustomerOrder struct {
	ID           uint    `json:"id"`
	TableInfoID  uint    `json:"table_id"`
	TableOrderID uint    `json:"table_order_id"`
	CustomerID   *uint   `json:"customer_id"`
	OrderNumber  string  `json:"order_number"`
	OrderAt      string  `json:"order_at"`
	Remark       string  `json:"remark"`
	Orders       []Order `json:"order"`
}
