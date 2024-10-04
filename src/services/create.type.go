package services

import "time"

type CreateTableOrderParams struct {
	Number        string
	ReceiptNumber string
	TableInfoID   uint
	CustomerID    *uint
	Status        string
	ReservedAt    *time.Time
	CancelAt      *time.Time
	OpenedAt      *time.Time
	CheckedOutAt  *time.Time
}
