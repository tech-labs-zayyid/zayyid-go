package model

type SalesPaymentResp struct {
	TransactionStatus string  `json:"transaction_status" db:"transaction_status"`
	TransactionId     string  `json:"transaction_id" db:"transaction_id"`
	StatusMessage     string  `json:"status_message" db:"status_message"`
	StatusCode        int     `json:"status_code" db:"status_code"`
	PaymentType       string  `json:"payment_type" db:"payment_type"`
	OrderId           string  `json:"order_id" db:"order_id"`
	GrossAmount       float64 `json:"gross_amount" db:"gross_amount"`
	FraudStatus       string  `json:"fraud_status" db:"fraud_status"`
	Bank              string  `json:"bank" db:"bank"`
	CreatedAt         string  `json:"created_at" db:"created_at"`
	UpdatedAt         string  `json:"updated_at" db:"updated_at"`
}
