package entity

import (
	"time"
)

type Payment struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID       int       `gorm:"not null;index" json:"order_id"`
	RedirectURL   string    `gorm:"type:text" json:"redirect_url,omitempty"`
	Subtotal      int       `gorm:"type:int" json:"subtotal,omitempty"`
	Status        string    `gorm:"type:varchar(50)" json:"status,omitempty"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	TransferType  string    `gorm:"type:varchar(50)" json:"transfer_type,omitempty"`
	TransactionID string    `gorm:"type:text" json:"transaction_id,omitempty"`
	Notification  string    `gorm:"type:text" json:"notification,omitempty"`
}
