package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetTodayTotalRevenue() (int, error) {
	var totalRevenue int
	err := repo.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0) 
		FROM transactions 
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&totalRevenue)
	if err != nil {
		return 0, err
	}
	return totalRevenue, nil
}

func (repo *ReportRepository) GetTodayTransactionCount() (int, error) {
	var count int
	err := repo.db.QueryRow(`
		SELECT COUNT(*) 
		FROM transactions 
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *ReportRepository) GetBestSellingProductToday() (*models.BestSellingProduct, error) {
	var productName string
	var totalQty int
	err := repo.db.QueryRow(`
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM products p
		LEFT JOIN transaction_details td ON p.id = td.product_id
		LEFT JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`).Scan(&productName, &totalQty)
	if err == sql.ErrNoRows {
		return &models.BestSellingProduct{
			Nama:       "N/A",
			QtyTerjual: 0,
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return &models.BestSellingProduct{
		Nama:       productName,
		QtyTerjual: totalQty,
	}, nil
}
