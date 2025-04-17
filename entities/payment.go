package entities

import "time"

type PaymentTransactionListScan struct {
	OrderId           string    `json:"order_id"`
	GrossAmount       int       `json:"gross_amount"`
	TotalAmount       int       `json:"total_amount"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
}

type PaymentTransactionListResponse struct {
	OrderId           string    `json:"order_id"`
	GrossAmount       int       `json:"gross_amount"`
	TotalAmount       int       `json:"total_amount"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
}
