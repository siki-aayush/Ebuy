package controllers

import (
	"e-buy/models"
	"e-buy/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type CustomerController struct {
	service *services.CustomerService
}

func NewCustomerController(service *services.CustomerService) *CustomerController {
	return &CustomerController{service: service}
}

func (c *CustomerController) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := c.service.CreateCustomer(r.Context(), customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (c *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	customer, err := c.service.GetCustomerByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func (c *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer.ID = id

	if err := c.service.UpdateCustomer(r.Context(), customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *CustomerController) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	if err := c.service.DeleteCustomer(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
