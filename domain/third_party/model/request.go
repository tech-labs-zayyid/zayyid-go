package model

type MidtransNotificationBodyReq struct {
	TransactionStatus string `json:"transaction_status"`
	TransactionId     string `json:"transaction_id"`
	StatusMessage     string `json:"status_message"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	PaymentType       string `json:"payment_type"`
	OrderId           string `json:"order_id"`
	GrossAmount       string `json:"gross_amount"`
	FraudStatus       string `json:"fraud_status"`
	Bank              string `json:"bank"`
}

type FrontendNotificationBodyReq struct {
	TransactionStatus string  `json:"transaction_status"`
	TransactionId     string  `json:"transaction_id"`
	StatusMessage     string  `json:"status_message"`
	StatusCode        int     `json:"status_code"`
	PaymentType       string  `json:"payment_type"`
	OrderId           string  `json:"order_id"`
	GrossAmount       float64 `json:"gross_amount"`
	FraudStatus       string  `json:"fraud_status"`
	Bank              string  `json:"bank"`
}
