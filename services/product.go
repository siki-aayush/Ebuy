package services

import (
	"context"
	"e-buy/models"
	"e-buy/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(ctx context.Context, product models.Product) (int, error) {
	// Simply call the repository function and return the result
	return s.productRepo.CreateProduct(ctx, product)
}

// GetProductByID retrieves a product by its ID
func (s *ProductService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	// Simply call the repository function and return the result
	return s.productRepo.GetProductByID(ctx, id)
}

// GetAllProducts retrieves all products from the database
func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	// Simply call the repository function and return the result
	return s.productRepo.GetAllProducts(ctx)
}

// UpdateProduct updates an existing product's details
func (s *ProductService) UpdateProduct(ctx context.Context, product models.Product) error {
	// Simply call the repository function and return the result
	return s.productRepo.UpdateProduct(ctx, product)
}

// DeleteProduct deletes a product by its ID
func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	// Simply call the repository function and return the result
	return s.productRepo.DeleteProduct(ctx, id)
}
