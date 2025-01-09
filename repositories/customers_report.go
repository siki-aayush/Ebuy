package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerReportRepository struct {
	db *pgxpool.Pool
}

func NewCustomerReportRepository(db *pgxpool.Pool) *CustomerReportRepository {
	return &CustomerReportRepository{db: db}
}

func (r *CustomerReportRepository) GetCustomerReport(ctx context.Context, startDate, endDate string) (map[string]interface{}, error) {
	query := `
    WITH filtered_customers AS (
        SELECT 
            c.id AS customer_id,
            c.signup_date,
            c.lifetime_value,
            COUNT(o.id) AS total_orders
        FROM customers c
        LEFT JOIN orders o ON c.id = o.customer_id
        WHERE 1 = 1
    `

	// Apply filters based on provided parameters
	if startDate != "" {
		query += fmt.Sprintf(`AND c.signup_date >= COALESCE('%s'::timestamp, '2000-01-01') `, startDate)
	}
	if endDate != "" {
		query += fmt.Sprintf(`AND c.signup_date <= COALESCE('%s'::timestamp, NOW()) `, endDate)
	}

	query += `
      GROUP BY c.id
  ),
  -- Aggregations
  aggregated_data as (
      SELECT 
      COUNT(customer_id) AS total_customers,
      AVG(total_orders) AS avg_order_frequency
      from filtered_customers
  )
  select * from aggregated_data
  `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	report := make(map[string]interface{})
	for rows.Next() {
		var totalCustomers sql.NullInt64
		var avgOrderFrequency sql.NullFloat64

		err = rows.Scan(&totalCustomers, &avgOrderFrequency)
		if err != nil {
			return nil, err
		}

		if totalCustomers.Valid {
			report["totalCustomers"] = totalCustomers.Int64
		} else {
			report["totalCustomers"] = 0
		}

		if avgOrderFrequency.Valid {
			report["avgOrderFrequency"] = avgOrderFrequency.Float64
		} else {
			report["avgOrderFrequency"] = nil
		}
	}

	return report, nil
}
