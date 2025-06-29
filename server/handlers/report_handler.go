package handlers

import (
	"net/http"
	"server/dto"
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

func (h *ReportHandler) GetSummary(c *gin.Context) {
	data, err := h.service.GetSummary()
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *ReportHandler) GetOrderReports(c *gin.Context) {
	var params dto.OrderReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}
	lists, pagination, err := h.service.GetOrderReports(params)

	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	// TODO : still hard coded, need to be dynamic. create a utils to handle this all at once
	if params.Export == "csv" {
		utils.ExportCSV(c, "orders_reports.csv", lists)
		return
	}

	if params.Export == "pdf" {
		utils.ExportPDF(c, "orders_reports.pdf", lists)
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": lists, "pagination": pagination})
}

func (h *ReportHandler) GetTicketSalesReports(c *gin.Context) {
	var params dto.TicketReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	lists, pagination, err := h.service.GetTicketSalesReports(params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	if params.Export == "csv" {
		utils.ExportCSV(c, "tickets_reports.csv", lists)
		return
	}

	if params.Export == "pdf" {
		utils.ExportPDF(c, "tickets_reports.pdf", lists)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tickets":    lists,
		"pagination": pagination,
	})
}

func (h *ReportHandler) GetPaymentReports(c *gin.Context) {
	var params dto.PaymentReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	lists, pagination, err := h.service.GetPaymentReports(params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	if params.Export == "csv" {
		utils.ExportCSV(c, "payments_reports.csv", lists)
		return
	}

	if params.Export == "pdf" {
		utils.ExportPDF(c, "payments_reports.pdf", lists)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"payments":   lists,
		"pagination": pagination,
	})
}

func (h *ReportHandler) GetRefundReports(c *gin.Context) {
	var params dto.RefundReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	lists, pagination, err := h.service.GetRefundReports(params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	if params.Export == "csv" {
		utils.ExportCSV(c, "refund_reports.csv", lists)
		return
	}

	if params.Export == "pdf" {
		utils.ExportPDF(c, "refund_reports.pdf", lists)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"refunds":    lists,
		"pagination": pagination,
	})
}

// handlers/report_handler.go
func (h *ReportHandler) GetWithdrawalReports(c *gin.Context) {
	var params dto.WithdrawalReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	lists, pagination, err := h.service.GetWithdrawalReports(params)
	if err != nil {
		utils.HandleServiceError(c, err, "failed to get withdrawal reports")
		return
	}

	if params.Export == "csv" {
		utils.ExportCSV(c, "withdrawal_reports.csv", lists)
		return
	}

	if params.Export == "pdf" {
		utils.ExportPDF(c, "withdrawal_reports.pdf", lists)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"withdrawals": lists,
		"pagination":  pagination,
	})
}
