package services

import (
	"context"
	"e-buy/models"
	"e-buy/repositories"
)

// OrderItemService represents the service for managing order items
type OrderItemService struct {
	repo *repositories.OrderItemRepository
}

// NewOrderItemService creates a new instance of OrderItemService
func NewOrderItemService(repo *repositories.OrderItemRepository) *OrderItemService {
	return &OrderItemService{repo: repo}
}

// CreateOrderItem creates a new order item
func (s *OrderItemService) CreateOrderItem(ctx context.Context, item models.OrderItem) error {
	return s.repo.CreateOrderItem(ctx, item)
}

// GetOrderItemsByOrderID retrieves order items by order ID
func (s *OrderItemService) GetOrderItemsByOrderID(ctx context.Context, orderID int) ([]models.OrderItem, error) {
	return s.repo.GetOrderItemsByOrderID(ctx, orderID)
}

// UpdateOrderItem updates an existing order item
func (s *OrderItemService) UpdateOrderItem(ctx context.Context, item models.OrderItem) error {
	return s.repo.UpdateOrderItem(ctx, item)
}

// DeleteOrderItem deletes an order item
func (s *OrderItemService) DeleteOrderItem(ctx context.Context, orderID, productID int) error {
	return s.repo.DeleteOrderItem(ctx, orderID, productID)
}
