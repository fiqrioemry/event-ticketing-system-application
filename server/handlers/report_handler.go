package handlers

import (
	"net/http"
	"server/services"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service services.ReportService
}

func NewReportHandler(service services.ReportService) *ReportHandler {
	return &ReportHandler{service}
}

func (h *ReportHandler) GetAllTransactions(c *gin.Context) {
	res, err := h.service.GetAllTransactions()
	if err != nil {
		utils.HandleServiceError(c, err, "failed to fetch transaction reports")
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ReportHandler) GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	res, err := h.service.GetTransactionByID(id)
	if err != nil {
		utils.HandleServiceError(c, err, "transaction not found")
		return
	}
	c.JSON(http.StatusOK, res)
}
