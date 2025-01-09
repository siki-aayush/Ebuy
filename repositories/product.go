package repositories

import (
	"context"
	"e-buy/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create
func (r *ProductRepository) CreateProduct(ctx context.Context, product models.Product) (int, error) {
	var id int
	query := `INSERT INTO products (name, category, price)
			  VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(ctx, query, product.Name, product.Category, product.Price).Scan(&id)
	return id, err
}

// Read
func (r *ProductRepository) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	product := &models.Product{}
	query := `SELECT id, name, category, price FROM products WHERE id=$1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&product.ID, &product.Name, &product.Category, &product.Price,
	)
	return product, err
}

func (r *ProductRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, category, price FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// Update
func (r *ProductRepository) UpdateProduct(ctx context.Context, product models.Product) error {
	query := `UPDATE products SET name=$1, category=$2, price=$3 WHERE id=$4`
	_, err := r.db.Exec(ctx, query, product.Name, product.Category, product.Price, product.ID)
	return err
}

// Delete
func (r *ProductRepository) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
