package models

import "time"

// Transaction represents a record in the transactions table.
type Transaction struct {
	ID            int       `json:"id"`
	OrderID       int       `json:"order_id"`
	PaymentStatus string    `json:"payment_status"` // 'SUCCESS' or 'FAILED'
	PaymentDate   time.Time `json:"payment_date"`
	TotalAmount   float64   `json:"total_amount"`
}
