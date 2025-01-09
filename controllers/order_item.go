package controllers

import (
	"e-buy/models"
	"e-buy/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type OrderItemController struct {
	service *services.OrderItemService
}

// NewOrderItemController creates a new instance of OrderItemController
func NewOrderItemController(service *services.OrderItemService) *OrderItemController {
	return &OrderItemController{service: service}
}

// CreateOrderItem creates a new order item
func (c *OrderItemController) CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	var item models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.service.CreateOrderItem(r.Context(), item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetOrderItems retrieves order items by order ID
func (c *OrderItemController) GetOrderItems(w http.ResponseWriter, r *http.Request) {
	orderIDParam := r.URL.Query().Get("order_id")
	orderID, _ := strconv.Atoi(orderIDParam)
	items, err := c.service.GetOrderItemsByOrderID(r.Context(), orderID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(items)
}

// UpdateOrderItem updates an existing order item
func (c *OrderItemController) UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	orderIDParam := r.URL.Query().Get("order_id")
	productIDParam := r.URL.Query().Get("product_id")
	orderID, _ := strconv.Atoi(orderIDParam)
	productID, _ := strconv.Atoi(productIDParam)

	var item models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.OrderID = orderID
	item.ProductID = productID

	if err := c.service.UpdateOrderItem(r.Context(), item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteOrderItem deletes an order item
func (c *OrderItemController) DeleteOrderItem(w http.ResponseWriter, r *http.Request) {
	orderIDParam := r.URL.Query().Get("order_id")
	productIDParam := r.URL.Query().Get("product_id")
	orderID, _ := strconv.Atoi(orderIDParam)
	productID, _ := strconv.Atoi(productIDParam)

	if err := c.service.DeleteOrderItem(r.Context(), orderID, productID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
