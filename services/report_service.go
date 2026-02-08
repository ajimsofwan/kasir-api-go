package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailyReport() (*models.DailyReport, error) {
	totalRevenue, err := s.repo.GetTodayTotalRevenue()
	if err != nil {
		return nil, err
	}

	totalTransactions, err := s.repo.GetTodayTransactionCount()
	if err != nil {
		return nil, err
	}

	bestSellingProduct, err := s.repo.GetBestSellingProductToday()
	if err != nil {
		return nil, err
	}

	return &models.DailyReport{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransactions,
		ProdukTerlaris: *bestSellingProduct,
	}, nil
}
