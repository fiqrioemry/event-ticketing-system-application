package services

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"

	"github.com/fiqrioemry/go-api-toolkit/response"
)

type AdminService interface {
	GetSummary() (*dto.SummaryReportResponse, error)
	GetAllEvents(params dto.EventQueryParams) ([]dto.AdminEventResponse, int, error)
	GetAllUsers(params dto.UserQueryParams) ([]dto.UserListResponse, int, error)
	GetOrderReports(params dto.OrderReportQueryParams) ([]dto.OrderReportResponse, int, error)
	GetRefundReports(params dto.RefundReportQueryParams) ([]dto.RefundReportResponse, int, error)
	GetPaymentReports(params dto.PaymentReportQueryParams) ([]dto.PaymentReportResponse, int, error)
	GetTicketSalesReports(params dto.TicketReportQueryParams) ([]dto.TicketSalesReportResponse, int, error)
	GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]dto.WithdrawalReportResponse, int, error)
}

type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo}
}

func (s *adminService) GetSummary() (*dto.SummaryReportResponse, error) {
	return s.repo.GetSummary()
}

func (s *adminService) GetAllEvents(params dto.EventQueryParams) ([]dto.AdminEventResponse, int, error) {
	// Get events from repository
	list, total, err := s.repo.GetAllEvents(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("Failed to get events", err)
	}

	var result []dto.AdminEventResponse
	for _, item := range list {
		// Initialize counters
		totalQuota := 0
		totalSold := 0
		startPrice := 0.0
		ticketCount := len(item.Tickets)
		revenue := 0.0

		// Calculate statistics from tickets
		for _, ticket := range item.Tickets {
			totalQuota += ticket.Quota
			totalSold += ticket.Sold
			revenue += float64(ticket.Sold) * ticket.Price

			// Set start price (lowest price) - ticket.Price is already float64
			if startPrice == 0.0 || ticket.Price < startPrice {
				startPrice = ticket.Price
			}
		}

		// Calculate remaining quota
		remainingQuota := totalQuota - totalSold

		// Determine availability status
		isAvailable := false
		if ticketCount > 0 {
			isAvailable = (item.Status == "active" || item.Status == "ongoing") && remainingQuota > 0
		}

		// Calculate completion percentage
		completionPercentage := 0.0
		if totalQuota > 0 {
			completionPercentage = (float64(totalSold) / float64(totalQuota)) * 100
		}

		// Build admin event response
		result = append(result, dto.AdminEventResponse{
			ID:          item.ID.String(),
			Image:       item.Image,
			Title:       item.Title,
			Description: item.Description,
			Location:    item.Location,
			StartPrice:  startPrice,
			IsAvailable: isAvailable,
			StartTime:   item.StartTime,
			EndTime:     item.EndTime,
			Status:      item.Status,
			Date:        item.Date,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,

			// Admin-specific fields
			TicketCount:          ticketCount,
			TotalQuota:           totalQuota,
			TotalSold:            totalSold,
			RemainingQuota:       remainingQuota,
			CompletionPercentage: completionPercentage,
			Revenue:              revenue,

			// Include ticket details
			Tickets: setTicketResponse(item.Tickets),
		})
	}

	return result, int(total), nil
}

// Helper function to map tickets to DTO with proper calculations
func setTicketResponse(tickets []models.Ticket) []dto.TicketResponse {
	var result []dto.TicketResponse
	for _, ticket := range tickets {

		result = append(result, dto.TicketResponse{
			ID:            ticket.ID.String(),
			EventID:       ticket.EventID.String(),
			Name:          ticket.Name,
			Price:         ticket.Price, // Already float64, no casting needed
			Limit:         ticket.Limit,
			Quota:         ticket.Quota,
			Sold:          ticket.Sold,
			Refundable:    ticket.Refundable,
			RefundPercent: ticket.RefundPercent,
		})
	}
	return result
}

func (s *adminService) GetAllUsers(params dto.UserQueryParams) ([]dto.UserListResponse, int, error) {
	users, total, err := s.repo.GetAllUsers(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to fetch user list", err)
	}

	var results []dto.UserListResponse
	for _, u := range users {
		results = append(results, dto.UserListResponse{
			ID:       u.ID.String(),
			Email:    u.Email,
			Role:     u.Role,
			Avatar:   u.Avatar,
			Fullname: u.Fullname,
			JoinedAt: u.CreatedAt,
		})
	}

	return results, int(total), nil
}

func (s *adminService) GetOrderReports(params dto.OrderReportQueryParams) ([]dto.OrderReportResponse, int, error) {
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

func (s *adminService) GetTicketSalesReports(params dto.TicketReportQueryParams) ([]dto.TicketSalesReportResponse, int, error) {

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

func (s *adminService) GetPaymentReports(params dto.PaymentReportQueryParams) ([]dto.PaymentReportResponse, int, error) {
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

func (s *adminService) GetRefundReports(params dto.RefundReportQueryParams) ([]dto.RefundReportResponse, int, error) {
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

func (s *adminService) GetWithdrawalReports(params dto.WithdrawalReportQueryParams) ([]dto.WithdrawalReportResponse, int, error) {
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
