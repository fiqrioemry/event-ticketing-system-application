package services

import (
	"server/dto"
	customErr "server/errors"
	"server/repositories"
	"server/utils"
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
		return nil, nil, customErr.NewInternalServerError("failed to retrieve order reports", err)
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
			CreatedAt:  order.CreatedAt,
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return reports, pagination, nil
}

func (s *reportService) GetTicketSalesReports(params dto.TicketReportQueryParams) ([]dto.TicketSalesReportResponse, *dto.PaginationResponse, error) {

	list, total, err := s.repo.GetTicketSalesReports(params)
	if err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to retrieve ticket sales reports", err)
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
		return nil, nil, customErr.NewInternalServerError("failed to retrieve payment reports", err)
	}

	var results []dto.PaymentReportResponse
	for _, p := range list {
		results = append(results, dto.PaymentReportResponse{
			PaymentID: p.ID.String(),
			OrderID:   p.OrderID.String(),
			Fullname:  p.Fullname,
			Email:     p.Email,
			Method:    p.Method,
			Amount:    p.Amount,
			Status:    p.Status,
			PaidAt:    p.PaidAt,
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil
}

func (s *reportService) GetRefundReports(params dto.RefundReportQueryParams) ([]dto.RefundReportResponse, *dto.PaginationResponse, error) {
	list, total, err := s.repo.GetRefundReports(params)
	if err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to retrieve refund reports", err)
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
			RefundedAt:   o.RefundedAt,
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return result, pagination, nil
}

// services/report_service.go
func (s *reportService) GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]dto.WithdrawalReportResponse, *dto.PaginationResponse, error) {
	list, total, err := s.repo.GetWithdrawalReports(params)
	if err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to retrieve withdrawal reports", err)
	}

	var result []dto.WithdrawalReportResponse
	for _, w := range list {

		result = append(result, dto.WithdrawalReportResponse{
			WithdrawalID: w.ID.String(),
			UserID:       w.UserID.String(),
			Fullname:     w.User.Fullname,
			Email:        w.User.Email,
			Amount:       w.Amount,
			Status:       w.Status,
			Reason:       w.Reason,
			CreatedAt:    w.CreatedAt,
			ApprovedAt:   w.ApprovedAt,
		})
	}

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return result, pagination, nil
}
