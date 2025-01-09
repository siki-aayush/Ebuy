package repositories

import (
	"context"
	"e-buy/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// Create
func (r *CustomerRepository) CreateCustomer(ctx context.Context, customer models.Customer) (int, error) {
	var id int
	query := `INSERT INTO customers (name, email, signup_date, location, lifetime_value)
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(ctx, query, customer.Name, customer.Email, customer.SignupDate, customer.Location, customer.LifetimeValue).Scan(&id)
	return id, err
}

// Read
func (r *CustomerRepository) GetCustomerByID(ctx context.Context, id int) (*models.Customer, error) {
	customer := &models.Customer{}
	query := `SELECT id, name, email, signup_date, location, lifetime_value FROM customers WHERE id=$1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&customer.ID, &customer.Name, &customer.Email, &customer.SignupDate, &customer.Location, &customer.LifetimeValue,
	)
	return customer, err
}

// Update
func (r *CustomerRepository) UpdateCustomer(ctx context.Context, customer models.Customer) error {
	query := `UPDATE customers SET name=$1, email=$2, signup_date=$3, location=$4, lifetime_value=$5 WHERE id=$6`
	_, err := r.db.Exec(ctx, query, customer.Name, customer.Email, customer.SignupDate, customer.Location, customer.LifetimeValue, customer.ID)
	return err
}

// Delete
func (r *CustomerRepository) DeleteCustomer(ctx context.Context, id int) error {
	query := `DELETE FROM customers WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
