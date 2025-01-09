package repositories

import (
	"context"
	"e-buy/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// OrderItemRepository represents the repository for managing order items
type OrderItemRepository struct {
	db *pgxpool.Pool
}

// NewOrderItemRepository creates a new instance of OrderItemRepository
func NewOrderItemRepository(db *pgxpool.Pool) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}

// Create inserts a new order item into the database
func (r *OrderItemRepository) CreateOrderItem(ctx context.Context, item models.OrderItem) error {
	query := `INSERT INTO order_items (order_id, product_id, quantity, price) 
			  VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, item.OrderID, item.ProductID, item.Quantity, item.Price)
	return err
}

// GetByOrderID retrieves all order items by order ID
func (r *OrderItemRepository) GetOrderItemsByOrderID(ctx context.Context, orderID int) ([]models.OrderItem, error) {
	var items []models.OrderItem
	query := `SELECT order_id, product_id, quantity, price FROM order_items WHERE order_id=$1`
	rows, err := r.db.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.OrderItem
		if err := rows.Scan(&item.OrderID, &item.ProductID, &item.Quantity, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Update updates an existing order item
func (r *OrderItemRepository) UpdateOrderItem(ctx context.Context, item models.OrderItem) error {
	query := `UPDATE order_items SET quantity=$1, price=$2 WHERE order_id=$3 AND product_id=$4`
	_, err := r.db.Exec(ctx, query, item.Quantity, item.Price, item.OrderID, item.ProductID)
	return err
}

// Delete removes an order item from the database
func (r *OrderItemRepository) DeleteOrderItem(ctx context.Context, orderID, productID int) error {
	query := `DELETE FROM order_items WHERE order_id=$1 AND product_id=$2`
	_, err := r.db.Exec(ctx, query, orderID, productID)
	return err
}
