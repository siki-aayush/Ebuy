package services

import (
	"context"
	"e-buy/repositories"
)

type CustomerReportService struct {
	repo *repositories.CustomerReportRepository
}

func NewCustomerReportService(repo *repositories.CustomerReportRepository) *CustomerReportService {
	return &CustomerReportService{repo: repo}
}

func (s *CustomerReportService) GenerateCustomerReport(ctx context.Context, startDate, endDate string) (map[string]interface{}, error) {
	return s.repo.GetCustomerReport(ctx, startDate, endDate)
}
