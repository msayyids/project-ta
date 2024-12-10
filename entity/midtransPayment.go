package entity

import "time"

type Payment struct {
	ID             int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	OrderID        int       `gorm:"not null;column:order_id" json:"order_id"`
	RedirectURL    string    `gorm:"column:redirect_url" json:"redirect_url"`
	GrossAmount    int       `gorm:"column:subtotal" json:"Gross_amount"`
	Status         string    `gorm:"column:status" json:"status"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	TransferType   string    `gorm:"column:transfer_type" json:"transfer_type"`
	Transaction_Id string    `gorm:"column:transaction_id" json:"transaction_id"`
	Notification   string    `gorm:"column:notification" json:"notification"`
}
type PaymentNotification struct {
	StatusCode        string  `json:"status_code"`
	TransactionID     string  `json:"transaction_id"`
	GrossAmount       float64 `json:"gross_amount"`
	Currency          string  `json:"currency"`
	OrderID           string  `json:"order_id"`
	PaymentType       string  `json:"payment_type"`
	SignatureKey      string  `json:"signature_key"`
	TransactionStatus string  `json:"transaction_status"`
	FraudStatus       string  `json:"fraud_status"`
	StatusMessage     string  `json:"status_message"`
	MerchantID        string  `json:"merchant_id"`
	TransactionTime   string  `json:"transaction_time"`
	ExpiryTime        string  `json:"expiry_time"`
}

// TableName untuk mengubah nama tabel di database jika diperlukan
func (Payment) TableName() string {
	return "payments"
}
