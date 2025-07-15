package services

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"

	"github.com/fiqrioemry/go-api-toolkit/response"
)

type ReportService interface {
	GetSummary() (*dto.SummaryReportResponse, error)
	GetOrderReports(params dto.OrderReportQueryParams) ([]dto.OrderReportResponse, int, error)
	GetRefundReports(params dto.RefundReportQueryParams) ([]dto.RefundReportResponse, int, error)
	GetPaymentReports(params dto.PaymentReportQueryParams) ([]dto.PaymentReportResponse, int, error)
	GetTicketSalesReports(params dto.TicketReportQueryParams) ([]dto.TicketSalesReportResponse, int, error)
	GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]dto.WithdrawalReportResponse, int, error)
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

func (s *reportService) GetOrderReports(params dto.OrderReportQueryParams) ([]dto.OrderReportResponse, int, error) {
	orders, total, err := s.repo.GetOrderReports(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to retrieve order reports", err)
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

	return reports, int(total), nil
}

func (s *reportService) GetTicketSalesReports(params dto.TicketReportQueryParams) ([]dto.TicketSalesReportResponse, int, error) {

	list, total, err := s.repo.GetTicketSalesReports(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to retrieve ticket sales reports", err)
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

	return reports, int(total), nil
}

func (s *reportService) GetPaymentReports(params dto.PaymentReportQueryParams) ([]dto.PaymentReportResponse, int, error) {
	list, total, err := s.repo.GetPaymentReports(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to retrieve payment reports", err)
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

	return results, int(total), nil
}

func (s *reportService) GetRefundReports(params dto.RefundReportQueryParams) ([]dto.RefundReportResponse, int, error) {
	list, total, err := s.repo.GetRefundReports(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to retrieve refund reports", err)
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

	return result, int(total), nil
}

// services/report_service.go
func (s *reportService) GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]dto.WithdrawalReportResponse, int, error) {
	list, total, err := s.repo.GetWithdrawalReports(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to retrieve withdrawal reports", err)
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

	return result, int(total), nil
}
