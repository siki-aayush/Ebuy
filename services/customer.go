package services

import (
	"context"
	"e-buy/models"
	"e-buy/repositories"
)

type CustomerService struct {
	repo *repositories.CustomerRepository
}

func NewCustomerService(repo *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, customer models.Customer) (int, error) {
	return s.repo.CreateCustomer(ctx, customer)
}

func (s *CustomerService) GetCustomerByID(ctx context.Context, id int) (*models.Customer, error) {
	return s.repo.GetCustomerByID(ctx, id)
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, customer models.Customer) error {
	return s.repo.UpdateCustomer(ctx, customer)
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id int) error {
	return s.repo.DeleteCustomer(ctx, id)
}
