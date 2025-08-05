package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/pagination"
	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	service services.AdminService
}

func NewAdminHandler(service services.AdminService) *AdminHandler {
	return &AdminHandler{service}
}

func (h *AdminHandler) GetSummary(c *gin.Context) {
	// fetch summary data
	data, err := h.service.GetSummary()
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Summary retrieved successfully", data)
}

func (h *AdminHandler) GetAllEvents(c *gin.Context) {
	var params dto.EventQueryParams
	// bind query params
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch events data for admin
	data, total, err := h.service.GetAllEvents(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// build pagination meta
	pag := pagination.Build(params.Page, params.Limit, total)

	response.OKWithPagination(c, "events retrieved successfully", data, pag)
}

func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	var params dto.UserQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	users, total, err := h.service.GetAllUsers(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	pagination := pagination.Build(params.Page, params.Limit, total)

	response.OKWithPagination(c, "Users retrieved successfully", users, pagination)
}

func (h *AdminHandler) GetOrderReports(c *gin.Context) {
	// bind query params
	var params dto.OrderReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch order reports
	lists, total, err := h.service.GetOrderReports(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// export as CSV
	if params.Export == "csv" {
		utils.ExportCSV(c, "orders_reports.csv", lists)
		return
	}

	// export as PDF
	if params.Export == "pdf" {
		utils.ExportPDF(c, "orders_reports.pdf", lists)
		return
	}

	// build pagination meta
	paginate := pagination.Build(params.Page, params.Limit, total)
	response.OKWithPagination(c, "Orders retrieved successfully", lists, paginate)
}

func (h *AdminHandler) GetTicketSalesReports(c *gin.Context) {
	// bind query params
	var params dto.TicketReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch ticket reports
	lists, total, err := h.service.GetTicketSalesReports(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// export as CSV
	if params.Export == "csv" {
		utils.ExportCSV(c, "tickets_reports.csv", lists)
		return
	}

	// export as PDF
	if params.Export == "pdf" {
		utils.ExportPDF(c, "tickets_reports.pdf", lists)
		return
	}

	// build pagination meta
	paginate := pagination.Build(params.Page, params.Limit, total)
	response.OKWithPagination(c, "Ticket sales reports retrieved successfully", lists, paginate)
}

func (h *AdminHandler) GetPaymentReports(c *gin.Context) {
	// bind query params
	var params dto.PaymentReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch payment reports
	lists, total, err := h.service.GetPaymentReports(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// export as CSV
	if params.Export == "csv" {
		utils.ExportCSV(c, "payments_reports.csv", lists)
		return
	}

	// export as PDF
	if params.Export == "pdf" {
		utils.ExportPDF(c, "payments_reports.pdf", lists)
		return
	}

	// build pagination meta
	paginate := pagination.Build(params.Page, params.Limit, total)
	response.OKWithPagination(c, "Payment reports retrieved successfully", lists, paginate)
}

func (h *AdminHandler) GetRefundReports(c *gin.Context) {
	// bind query params
	var params dto.RefundReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch refund reports
	lists, total, err := h.service.GetRefundReports(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// export as CSV
	if params.Export == "csv" {
		utils.ExportCSV(c, "refund_reports.csv", lists)
		return
	}

	// export as PDF
	if params.Export == "pdf" {
		utils.ExportPDF(c, "refund_reports.pdf", lists)
		return
	}

	// build pagination meta
	paginate := pagination.Build(params.Page, params.Limit, total)
	response.OKWithPagination(c, "Refund reports retrieved successfully", lists, paginate)
}

func (h *AdminHandler) GetWithdrawalReports(c *gin.Context) {
	// bind query params
	var params dto.WithdrawalReportQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch withdrawal reports
	lists, total, err := h.service.GetWithdrawalReports(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// export as CSV
	if params.Export == "csv" {
		utils.ExportCSV(c, "withdrawal_reports.csv", lists)
		return
	}

	// export as PDF
	if params.Export == "pdf" {
		utils.ExportPDF(c, "withdrawal_reports.pdf", lists)
		return
	}

	// build pagination meta
	paginate := pagination.Build(params.Page, params.Limit, total)
	response.OKWithPagination(c, "withdrawal reports retrieved successfully", lists, paginate)
}
