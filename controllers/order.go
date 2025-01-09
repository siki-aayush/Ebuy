package controllers

import (
	"e-buy/models"
	"e-buy/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

// CreateOrder creates a new order
func (o *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := o.service.CreateOrder(r.Context(), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// GetOrder retrieves an order by ID or all orders if no ID is provided
func (o *OrderController) GetOrder(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")

	if idParam != "" { // Get order by ID
		id, _ := strconv.Atoi(idParam)
		order, err := o.service.GetOrderByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(order)
	} else { // Get all orders
		orders, err := o.service.GetAllOrders(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(orders)
	}
}

// UpdateOrder updates an existing order
func (o *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order.ID = id

	if err := o.service.UpdateOrder(r.Context(), order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteOrder deletes an order by ID
func (o *OrderController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	if err := o.service.DeleteOrder(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
