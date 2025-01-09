package main

import (
	"context"
	"e-buy/controllers"
	"e-buy/database"
	"e-buy/repositories"
	"e-buy/services"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	godotenv.Load()

	// Setup database connection
	conPool, err := pgxpool.NewWithConfig(context.Background(), database.Config())

	if err != nil {
		log.Fatal("Error creating connection to the database: ", err)
	}

	connection, err := conPool.Acquire(context.Background())

	if err != nil {
		log.Fatal("Error acquiring connection from the database pool: ", err)
	}

	defer connection.Release()

	// Initialize customer components
	customerRepo := repositories.NewCustomerRepository(conPool)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	// Define routes
	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			customerController.CreateCustomer(w, r)
		case http.MethodGet:
			customerController.GetCustomer(w, r)
		case http.MethodPut:
			customerController.UpdateCustomer(w, r)
		case http.MethodDelete:
			customerController.DeleteCustomer(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Initialize products components
	productRepo := repositories.NewProductRepository(conPool)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			productController.CreateProduct(w, r)
		case http.MethodGet:
			productController.GetProduct(w, r)
		case http.MethodPut:
			productController.UpdateProduct(w, r)
		case http.MethodDelete:
			productController.DeleteProduct(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Initialize order components
	orderRepo := repositories.NewOrderRepository(conPool)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderController.CreateOrder(w, r)
		case http.MethodGet:
			orderController.GetOrder(w, r)
		case http.MethodPut:
			orderController.UpdateOrder(w, r)
		case http.MethodDelete:
			orderController.DeleteOrder(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Initialize order items components
	orderItemRepo := repositories.NewOrderItemRepository(conPool)
	orderItemService := services.NewOrderItemService(orderItemRepo)
	orderItemController := controllers.NewOrderItemController(orderItemService)

	// Register routes
	http.HandleFunc("/order_items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			orderItemController.CreateOrderItem(w, r)
		case http.MethodGet:
			orderItemController.GetOrderItems(w, r)
		case http.MethodPut:
			orderItemController.UpdateOrderItem(w, r)
		case http.MethodDelete:
			orderItemController.DeleteOrderItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Initialize transactions components
	transactionRepo := repositories.NewTransactionRepository(conPool)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	// Register routes
	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			transactionController.CreateTransaction(w, r)
		case http.MethodGet:
			transactionController.GetTransactionsByOrderID(w, r)
		case http.MethodPut:
			transactionController.UpdateTransaction(w, r)
		case http.MethodDelete:
			transactionController.DeleteTransaction(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	reportRepo := repositories.NewReportRepository(conPool)
	reportService := services.NewReportService(reportRepo)
	reportController := controllers.NewReportController(reportService)

	http.HandleFunc("/reports/sales", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		reportController.GetSalesReport(w, r)
	})

	customerReportRepo := repositories.NewCustomerReportRepository(conPool)
	customerReportService := services.NewCustomerReportService(customerReportRepo)
	customerReportController := controllers.NewCustomerReportController(customerReportService)
	http.HandleFunc("/reports/customers", customerReportController.GetCustomerReport)

	// Start the server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
