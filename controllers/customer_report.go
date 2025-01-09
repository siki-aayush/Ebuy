package controllers

import (
	"e-buy/services"
	"encoding/json"
	"net/http"
)

type CustomerReportController struct {
	service *services.CustomerReportService
}

func NewCustomerReportController(service *services.CustomerReportService) *CustomerReportController {
	return &CustomerReportController{service: service}
}

func (c *CustomerReportController) GetCustomerReport(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	startDate := query.Get("startDate")
	endDate := query.Get("endDate")

	report, err := c.service.GenerateCustomerReport(r.Context(), startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
