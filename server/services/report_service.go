package services

import (
	"server/dto"
	"server/repositories"
	"server/utils"
	"time"
)

type ReportService interface {
	GetSummary() (*dto.SummaryReportResponse, error)
	GetOrderReports(params dto.OrderReportQueryParams) ([]dto.OrderReportResponse, *dto.PaginationResponse, error)
	GetTicketSalesReports(params dto.TicketReportQueryParams) ([]dto.TicketSalesReportResponse, *dto.PaginationResponse, error)
	GetPaymentReports(params dto.PaymentReportQueryParams) ([]dto.PaymentReportResponse, *dto.PaginationResponse, error)
	GetRefundReports(params dto.RefundReportQueryParams) ([]dto.RefundReportResponse, *dto.PaginationResponse, error)
	GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]dto.WithdrawalReportResponse, *dto.PaginationResponse, error)
}

type reportService struct {
	repo repositories.ReportRepository
}

func NewReportService(repo repositories.ReportRepository) ReportService {
	return &reportService{repo}
}

func (s *reportService) GetSummary() (*dto.SummaryReportResponse, error) {
	return s.repo.GetSummary()
}

func (s *reportService) GetOrderReports(params dto.OrderReportQueryParams) ([]dto.OrderReportResponse, *dto.PaginationResponse, error) {
	orders, total, err := s.repo.GetOrderReports(params)
	if err != nil {
		return nil, nil, err
	}

	var reports []dto.OrderReportResponse
	for _, order := range orders {
		reports = append(reports, dto.OrderReportResponse{
			OrderID:    order.ID.String(),
			Fullname:   order.Fullname,
			Email:      order.Email,
			EventTitle: order.Event.Title,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			CreatedAt:  order.CreatedAt.Format(time.RFC3339),
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return reports, pagination, nil
}

func (s *reportService) GetTicketSalesReports(params dto.TicketReportQueryParams) ([]dto.TicketSalesReportResponse, *dto.PaginationResponse, error) {

	list, total, err := s.repo.GetTicketSalesReports(params)
	if err != nil {
		return nil, nil, err
	}

	var reports []dto.TicketSalesReportResponse
	for _, ticket := range list {
		reports = append(reports, dto.TicketSalesReportResponse{
			EventTitle:  ticket.Event.Title,
			TicketName:  ticket.Name,
			TicketPrice: ticket.Price,
			Quota:       ticket.Quota,
			Sold:        ticket.Sold,
			Remaining:   ticket.Quota - ticket.Sold,
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return reports, pagination, nil
}

func (s *reportService) GetPaymentReports(params dto.PaymentReportQueryParams) ([]dto.PaymentReportResponse, *dto.PaginationResponse, error) {
	list, total, err := s.repo.GetPaymentReports(params)
	if err != nil {
		return nil, nil, err
	}

	var results []dto.PaymentReportResponse
	for _, p := range list {
		var paidAtStr *string
		if p.PaidAt != nil {
			str := p.PaidAt.Format(time.RFC3339)
			paidAtStr = &str
		}

		results = append(results, dto.PaymentReportResponse{
			PaymentID: p.ID.String(),
			OrderID:   p.OrderID.String(),
			Fullname:  p.Fullname,
			Email:     p.Email,
			Method:    p.Method,
			Amount:    p.Amount,
			Status:    p.Status,
			PaidAt:    paidAtStr,
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil
}

func (s *reportService) GetRefundReports(params dto.RefundReportQueryParams) ([]dto.RefundReportResponse, *dto.PaginationResponse, error) {
	list, total, err := s.repo.GetRefundReports(params)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.RefundReportResponse
	for _, o := range list {
		result = append(result, dto.RefundReportResponse{
			OrderID:      o.ID.String(),
			Fullname:     o.Fullname,
			Email:        o.Email,
			EventTitle:   o.Event.Title,
			RefundAmount: o.RefundAmount,
			RefundReason: o.RefundReason,
			RefundedAt:   o.RefundedAt.Format(time.RFC3339),
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return result, pagination, nil
}

// services/report_service.go
func (s *reportService) GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]dto.WithdrawalReportResponse, *dto.PaginationResponse, error) {
	list, total, err := s.repo.GetWithdrawalReports(params)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.WithdrawalReportResponse
	for _, w := range list {
		var approvedAt *string
		if w.ApprovedAt != nil {
			str := w.ApprovedAt.Format(time.RFC3339)
			approvedAt = &str
		}

		result = append(result, dto.WithdrawalReportResponse{
			WithdrawalID: w.ID.String(),
			UserID:       w.UserID.String(),
			Fullname:     w.User.Fullname,
			Email:        w.User.Email,
			Amount:       w.Amount,
			Status:       w.Status,
			Reason:       w.Reason,
			CreatedAt:    w.CreatedAt.Format(time.RFC3339),
			ApprovedAt:   approvedAt,
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return result, pagination, nil
}
