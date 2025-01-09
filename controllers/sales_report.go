package controllers

import (
	"e-buy/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReportController struct {
	service *services.ReportService
}

func NewReportController(service *services.ReportService) *ReportController {
	return &ReportController{service: service}
}

func (c *ReportController) GetSalesReport(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")
	category := query.Get("category")
	location := query.Get("location")
	productIDParam := query.Get("productID")

	var productID int
	var err error

	if productIDParam != "" {
		productID, err = strconv.Atoi(productIDParam)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
	}

	report, err := c.service.GenerateSalesReport(r.Context(), startDate, endDate, category, location, productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
