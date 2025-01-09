package controllers

import (
	"e-buy/models"
	"e-buy/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTransactionController(service *services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// Create
func (c *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.service.CreateTransaction(r.Context(), transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Get by ID
func (c *TransactionController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	transaction, err := c.service.GetTransactionByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(transaction)
}

// Get by OrderID
func (c *TransactionController) GetTransactionsByOrderID(w http.ResponseWriter, r *http.Request) {
	orderIDParam := r.URL.Query().Get("order_id")
	orderID, _ := strconv.Atoi(orderIDParam)
	transactions, err := c.service.GetTransactionsByOrderID(r.Context(), orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(transactions)
}

// Update
func (c *TransactionController) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction.ID = id

	if err := c.service.UpdateTransaction(r.Context(), transaction); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete
func (c *TransactionController) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	if err := c.service.DeleteTransaction(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
