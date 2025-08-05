package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/services"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"
	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type WithdrawalHandler struct {
	service    services.WithdrawalService
	repository repositories.AuditLogRepository
}

func NewWithdrawalHandler(service services.WithdrawalService, repository repositories.AuditLogRepository) *WithdrawalHandler {
	return &WithdrawalHandler{service, repository}
}

func (h *WithdrawalHandler) CreateWithdrawal(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var req dto.CreateWithdrawalRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	res, err := h.service.CreateWithdrawal(userID, req)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "Withdrawal request created successfully", res)
}

func (h *WithdrawalHandler) GetAllWithdrawals(c *gin.Context) {
	res, err := h.service.GetAllWithdrawals()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Withdrawals retrieved successfully", res)
}

func (h *WithdrawalHandler) ReviewWithdrawal(c *gin.Context) {
	adminID := utils.MustGetUserID(c)
	id := c.Param("id")
	var body struct {
		Status string `json:"status" binding:"required,oneof=approved rejected"`
	}

	if !utils.BindAndValidateJSON(c, &body) {
		return
	}
	res, err := h.service.ReviewWithdrawal(id, adminID, body.Status)
	if err != nil {
		response.Error(c, err)
		return
	}

	auditLog := utils.BuildAuditLog(c, adminID, "review", "withdrawal", res)

	go h.repository.Create(c.Request.Context(), auditLog)

	response.OK(c, "Withdrawal reviewed successfully", res)
}
