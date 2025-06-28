package services

import (
	"server/dto"
	"server/repositories"
	"time"
)

type ReportService interface {
	GetAllTransactions() ([]dto.TransactionReportResponseOutput, error)
	GetTransactionByID(id string) (*dto.TransactionReportResponseOutput, error)
}

type reportService struct {
	repo repositories.ReportRepository
}

func NewReportService(repo repositories.ReportRepository) ReportService {
	return &reportService{repo}
}

func (s *reportService) GetAllTransactions() ([]dto.TransactionReportResponseOutput, error) {
	records, err := s.repo.GetAllTransactions()
	if err != nil {
		return nil, err
	}

	var result []dto.TransactionReportResponseOutput
	for _, r := range records {
		result = append(result, formatReportDTO(r))
	}
	return result, nil
}

func (s *reportService) GetTransactionByID(id string) (*dto.TransactionReportResponseOutput, error) {
	r, err := s.repo.GetTransactionByID(id)
	if err != nil {
		return nil, err
	}
	res := formatReportDTO(*r)
	return &res, nil
}

func formatReportDTO(r dto.TransactionReportResponse) dto.TransactionReportResponseOutput {
	var paidAtStr *string
	if r.PaidAt != nil {
		str := r.PaidAt.Format(time.RFC3339)
		paidAtStr = &str
	}

	var refundedAtStr *string
	if r.RefundedAt != nil {
		str := r.RefundedAt.Format(time.RFC3339)
		refundedAtStr = &str
	}

	netRevenue := r.TotalPaid
	if r.RefundAmount != nil {
		netRevenue -= *r.RefundAmount
	}

	return dto.TransactionReportResponseOutput{
		OrderID:      r.OrderID,
		UserName:     r.UserName,
		UserEmail:    r.UserEmail,
		EventTitle:   r.EventTitle,
		TotalPaid:    r.TotalPaid,
		RefundAmount: r.RefundAmount,
		NetRevenue:   netRevenue,
		Status:       r.Status,
		Method:       r.Method,
		PaidAt:       paidAtStr,
		RefundedAt:   refundedAtStr,
	}
}
