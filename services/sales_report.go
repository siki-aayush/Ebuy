package services

import (
	"context"
	"e-buy/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GenerateSalesReport(ctx context.Context, startDate, endDate, category, location string, productID int) (map[string]interface{}, error) {
	return s.repo.GetSalesReport(ctx, startDate, endDate, category, location, productID)
}
