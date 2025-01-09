package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReportRepository struct {
	db *pgxpool.Pool
}

func NewReportRepository(db *pgxpool.Pool) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetSalesReport(ctx context.Context, startDate, endDate, category, location string, productID int) (map[string]interface{}, error) {
	query := `
    WITH filtered_orders AS (
      SELECT
          o.id AS order_id,
          o.order_date,
          o.customer_id,
          c.location AS customer_location,
          oi.product_id,
          oi.quantity,
          p.category,
          oi.price AS total_product_sales
      FROM
          orders o
          JOIN order_items oi ON o.id = oi.order_id
          JOIN customers c ON o.customer_id = c.id
          JOIN products p ON oi.product_id = p.id
      WHERE 1 = 1
      `

	if startDate != "" {
		query += fmt.Sprintf(`AND o.order_date >= COALESCE('%s'::timestamp, '2000-01-01') `, startDate)
	}
	if endDate != "" {
		query += fmt.Sprintf(`AND o.order_date <= COALESCE('%s'::timestamp, NOW())`, endDate)
	}
	if category != "" {
		query += fmt.Sprintf(`AND p.category = coalesce('%s', p.category)`, category)
	}
	if location != "" {
		query += fmt.Sprintf(`AND c.location = COALESCE('%s', c.location)`, location)
	}
	if productID != 0 {
		query += fmt.Sprintf(`AND oi.product_id = COALESCE(%d, oi.product_id)`, productID)
	}

	query +=
		`),
    aggregated_data AS (
        SELECT
            SUM(total_product_sales) AS total_sales,
            AVG(total_product_sales) AS avg_order_value,
            SUM(quantity) AS num_products_sold
        FROM
            filtered_orders
    )
    SELECT * from aggregated_data
  `
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	report := make(map[string]interface{})
	for rows.Next() {
		var totalSales sql.NullFloat64
		var avgOrderValue sql.NullFloat64
		var numProductsSold sql.NullInt64

		err = rows.Scan(&totalSales, &avgOrderValue, &numProductsSold)

		if err != nil {
			return nil, err
		}

		if totalSales.Valid {
			report["totalSales"] = totalSales.Float64
		} else {
			report["totalSales"] = 0
		}

		if avgOrderValue.Valid {
			report["averageOrderValue"] = avgOrderValue.Float64
		} else {
			report["averageOrderValue"] = 0
		}

		if numProductsSold.Valid {
			report["numberOfProductsSold"] = numProductsSold.Int64
		} else {
			report["numberOfProductsSold"] = 0
		}

	}
	return report, nil
}
