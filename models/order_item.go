package models

// OrderItem represents an item in an order
type OrderItem struct {
	OrderID   int     `json:"order_id"`   // The ID of the order
	ProductID int     `json:"product_id"` // The ID of the product
	Quantity  int     `json:"quantity"`   // The quantity of the product in the order
	Price     float64 `json:"price"`      // The price of the product at the time of order
}
