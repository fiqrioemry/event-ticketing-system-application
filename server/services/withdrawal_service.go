package services

import (
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/google/uuid"
)

type WithdrawalService interface {
	CreateWithdrawal(userID string, req dto.CreateWithdrawalRequest) (*dto.WithdrawalResponse, error)
	GetAllWithdrawals() ([]dto.WithdrawalResponse, error)
	ReviewWithdrawal(id, adminID, status string) (*dto.WithdrawalResponse, error)
}

type withdrawalService struct {
	repo repositories.WithdrawalRepository
}

func NewWithdrawalService(repo repositories.WithdrawalRepository) WithdrawalService {
	return &withdrawalService{repo}
}

func (s *withdrawalService) CreateWithdrawal(userID string, req dto.CreateWithdrawalRequest) (*dto.WithdrawalResponse, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, response.NewNotFound("user not found")
	}

	if user.Balance < req.Amount {
		return nil, response.NewBadRequest("insufficient balance")
	}

	withdrawal := &models.WithdrawalRequest{
		ID:        uuid.New(),
		UserID:    user.ID,
		Amount:    req.Amount,
		Status:    "pending",
		Reason:    req.Reason,
		CreatedAt: time.Now(),
	}
	if err := s.repo.CreateWithdrawal(withdrawal); err != nil {
		return nil, err
	}
	return toWithdrawalDTO(withdrawal), nil
}

func (s *withdrawalService) GetAllWithdrawals() ([]dto.WithdrawalResponse, error) {
	list, err := s.repo.GetAllWithdrawals()
	if err != nil {
		return nil, response.NewInternalServerError("failed to retrieve withdrawals", err)
	}

	var res []dto.WithdrawalResponse
	for _, w := range list {
		res = append(res, *toWithdrawalDTO(&w))
	}
	return res, nil
}

func (s *withdrawalService) ReviewWithdrawal(id, adminID, status string) (*dto.WithdrawalResponse, error) {
	w, err := s.repo.GetWithdrawalByID(id)
	if err != nil || w == nil {
		return nil, response.NewNotFound("withdrawal request not found").WithContext("withdrawalID", id)
	}

	if w.Status != "pending" {
		return nil, response.NewBadRequest("withdrawal already reviewed")
	}

	w.Status = status
	w.ApprovedAt = &time.Time{}
	*w.ApprovedAt = time.Now()

	if err := s.repo.UpdateWithdrawal(w); err != nil {
		return nil, err
	}

	if status == "approved" {
		_ = s.repo.DecreaseUserBalance(w.UserID.String(), w.Amount)
	}

	return toWithdrawalDTO(w), nil
}

func toWithdrawalDTO(w *models.WithdrawalRequest) *dto.WithdrawalResponse {
	res := &dto.WithdrawalResponse{
		ID:         w.ID.String(),
		UserID:     w.UserID.String(),
		Amount:     w.Amount,
		Status:     w.Status,
		Reason:     w.Reason,
		CreatedAt:  w.CreatedAt,
		ApprovedAt: w.ApprovedAt,
	}

	return res
}
