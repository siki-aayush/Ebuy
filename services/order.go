package services

import (
	"context"
	"e-buy/models"
	"e-buy/repositories"
)

type OrderService struct {
	repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(ctx context.Context, order models.Order) (int, error) {
	return s.repo.CreateOrder(ctx, order)
}

// GetOrderByID retrieves an order by ID
func (s *OrderService) GetOrderByID(ctx context.Context, id int) (*models.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

// GetAllOrders retrieves all orders
func (s *OrderService) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	return s.repo.GetAllOrders(ctx)
}

// UpdateOrder updates an existing order
func (s *OrderService) UpdateOrder(ctx context.Context, order models.Order) error {
	return s.repo.UpdateOrder(ctx, order)
}

// DeleteOrder deletes an order by ID
func (s *OrderService) DeleteOrder(ctx context.Context, id int) error {
	return s.repo.DeleteOrder(ctx, id)
}
