package repositories

import (
	"context"
	"e-buy/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create
func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction models.Transaction) (int, error) {
	var id int
	query := `INSERT INTO transactions (order_id, payment_status, payment_date, total_amount)
              VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(ctx, query, transaction.OrderID, transaction.PaymentStatus, transaction.PaymentDate, transaction.TotalAmount).Scan(&id)
	return id, err
}

// Get by ID
func (r *TransactionRepository) GetTransactionByID(ctx context.Context, id int) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	query := `SELECT id, order_id, payment_status, payment_date, total_amount FROM transactions WHERE id = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(
		&transaction.ID, &transaction.OrderID, &transaction.PaymentStatus, &transaction.PaymentDate, &transaction.TotalAmount,
	)
	return transaction, err
}

// Get by OrderID
func (r *TransactionRepository) GetTransactionsByOrderID(ctx context.Context, orderID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := `SELECT id, order_id, payment_status, payment_date, total_amount FROM transactions WHERE order_id = $1`
	rows, err := r.db.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.OrderID, &transaction.PaymentStatus, &transaction.PaymentDate, &transaction.TotalAmount); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// Update
func (r *TransactionRepository) UpdateTransaction(ctx context.Context, transaction models.Transaction) error {
	query := `UPDATE transactions SET payment_status=$1, payment_date=$2, total_amount=$3 WHERE id=$4`
	_, err := r.db.Exec(ctx, query, transaction.PaymentStatus, transaction.PaymentDate, transaction.TotalAmount, transaction.ID)
	return err
}

// Delete
func (r *TransactionRepository) DeleteTransaction(ctx context.Context, id int) error {
	query := `DELETE FROM transactions WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
