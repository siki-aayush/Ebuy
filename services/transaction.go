package services

import (
	"context"
	"e-buy/models"
	"e-buy/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// Create
func (s *TransactionService) CreateTransaction(ctx context.Context, transaction models.Transaction) (int, error) {
	return s.repo.CreateTransaction(ctx, transaction)
}

// Get by ID
func (s *TransactionService) GetTransactionByID(ctx context.Context, id int) (*models.Transaction, error) {
	return s.repo.GetTransactionByID(ctx, id)
}

// Get by OrderID
func (s *TransactionService) GetTransactionsByOrderID(ctx context.Context, orderID int) ([]models.Transaction, error) {
	return s.repo.GetTransactionsByOrderID(ctx, orderID)
}

// Update
func (s *TransactionService) UpdateTransaction(ctx context.Context, transaction models.Transaction) error {
	return s.repo.UpdateTransaction(ctx, transaction)
}

// Delete
func (s *TransactionService) DeleteTransaction(ctx context.Context, id int) error {
	return s.repo.DeleteTransaction(ctx, id)
}
