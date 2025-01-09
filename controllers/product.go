package controllers

import (
	"e-buy/models"
	"e-buy/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ProductController struct {
	service *services.ProductService
}

func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// CreateProduct creates a new product
func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := p.service.CreateProduct(r.Context(), product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// GetProduct retrieves a product by ID
func (p *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		product, err := p.service.GetAllProducts(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(product)
	} else {
		id, _ := strconv.Atoi(idParam)
		product, err := p.service.GetProductByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(product)
	}

}

// UpdateProduct updates an existing product by ID
func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product.ID = id

	if err := p.service.UpdateProduct(r.Context(), product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct deletes a product by ID
func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)
	if err := p.service.DeleteProduct(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
