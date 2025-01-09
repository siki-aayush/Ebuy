package repositories

import (
	"context"
	"e-buy/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

// CreateOrder creates a new order in the database
func (r *OrderRepository) CreateOrder(ctx context.Context, order models.Order) (int, error) {
	var id int
	query := `INSERT INTO orders (customer_id, order_date, status)
			  VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(ctx, query, order.CustomerID, order.OrderDate, order.Status).Scan(&id)
	return id, err
}

// GetOrderByID retrieves an order by its ID
func (r *OrderRepository) GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
	order := &models.Order{}
	query := `SELECT id, customer_id, order_date, status FROM orders WHERE id=$1`
	err := r.db.QueryRow(ctx, query, id).Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.Status)
	return order, err
}

// GetAllOrders retrieves all orders
func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	var orders []models.Order
	query := `SELECT id, customer_id, order_date, status FROM orders`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// UpdateOrder updates an existing order
func (r *OrderRepository) UpdateOrder(ctx context.Context, order models.Order) error {
	query := `UPDATE orders SET customer_id=$1, order_date=$2, status=$3 WHERE id=$4`
	_, err := r.db.Exec(ctx, query, order.CustomerID, order.OrderDate, order.Status, order.ID)
	return err
}

// DeleteOrder deletes an order by ID
func (r *OrderRepository) DeleteOrder(ctx context.Context, id int) error {
	query := `DELETE FROM orders WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
