package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
	"strings"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetDailyReport(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReportHandler) GetDailyReport(w http.ResponseWriter, r *http.Request) {
	// Handle /api/report/hari-ini
	if !strings.HasSuffix(r.URL.Path, "/hari-ini") {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	report, err := h.service.GetDailyReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
