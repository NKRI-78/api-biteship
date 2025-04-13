package entities

type PaymentTransactionListScan struct {
	OrderId           string `json:"order_id"`
	GrossAmount       int    `json:"gross_amount"`
	TotalAmount       int    `json:"total_amount"`
	TransactionStatus string `json:"transaction_status"`
}

type PaymentTransactionListResponse struct {
	OrderId           string `json:"order_id"`
	GrossAmount       int    `json:"gross_amount"`
	TotalAmount       int    `json:"total_amount"`
	TransactionStatus string `json:"transaction_status"`
}
