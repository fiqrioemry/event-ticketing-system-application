package handlers

import (
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type WithdrawalHandler struct {
	service services.WithdrawalService
}

func NewWithdrawalHandler(service services.WithdrawalService) *WithdrawalHandler {
	return &WithdrawalHandler{service}
}

func (h *WithdrawalHandler) CreateWithdrawal(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var req dto.CreateWithdrawalRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}
	res, err := h.service.CreateWithdrawal(userID, req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *WithdrawalHandler) GetAllWithdrawals(c *gin.Context) {
	res, err := h.service.GetAllWithdrawals()
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
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
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
