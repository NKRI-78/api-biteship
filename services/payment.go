package services

import (
	"database/sql"
	"errors"
	"superapps/entities"
	helper "superapps/helpers"
)

func TransactionListPayment() (map[string]any, error) {

	var transaction entities.PaymentTransactionListScan
	var dataTransaction = make([]entities.PaymentTransactionListResponse, 0)

	query := `SELECT orderId AS order_id, grossAmount AS gross_amount, totalAmount AS total_amount, transactionStatus AS transaction_status 
	FROM Payments`

	if dbPayment == nil {
		return nil, errors.New("❌ dbPayment connection is nil")
	}

	var rows *sql.Rows
	var err error

	rows, err = dbPayment.Debug().Raw(query).Rows()

	if err != nil {
		helper.Logger("error", "In Server: "+err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		errTransactionRows := dbPayment.ScanRows(rows, &transaction)

		if errTransactionRows != nil {
			helper.Logger("error", "In Server: "+errTransactionRows.Error())
			return nil, errors.New(errTransactionRows.Error())
		}

		dataTransaction = append(dataTransaction, entities.PaymentTransactionListResponse{
			OrderId:           transaction.OrderId,
			GrossAmount:       transaction.GrossAmount,
			TotalAmount:       transaction.TotalAmount,
			TransactionStatus: transaction.TransactionStatus,
		})
	}

	return map[string]any{
		"data": dataTransaction,
	}, nil
}
