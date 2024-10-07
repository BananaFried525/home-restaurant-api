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
