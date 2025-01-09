package models

import "time"

// Order represents an order in the system
type Order struct {
	ID         int       `json:"id"`          // The unique ID of the order
	CustomerID int       `json:"customer_id"` // The ID of the customer placing the order
	OrderDate  time.Time `json:"order_date"`  // The date and time when the order was placed
	Status     string    `json:"status"`      // The status of the order (e.g., PENDING, COMPLETED, CANCELED)
}
