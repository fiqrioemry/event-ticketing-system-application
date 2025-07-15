package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/pagination"
	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) CreateNewOrder(c *gin.Context) {
	// extract user ID
	userID := utils.MustGetUserID(c)

	// bind request data
	var req dto.CreateOrderRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	// create order record
	orderID, err := h.service.CreateNewOrder(req, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, "Order created successfully", orderID)
}

func (h *OrderHandler) GetMyOrders(c *gin.Context) {
	// extract user ID
	userID := utils.MustGetUserID(c)

	// bind query params
	var params dto.OrderQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	// apply pagination defaults
	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	// fetch user orders
	orders, total, err := h.service.GetMyOrders(userID, params)
	if err != nil {
		response.Error(c, err)
		return
	}

	// build pagination meta
	pag := pagination.Build(params.Page, params.Limit, total)

	response.OKWithPagination(c, "orders retrieved successfully", orders, pag)
}

func (h *OrderHandler) GetOrderDetail(c *gin.Context) {
	// extract order ID
	orderID := c.Param("id")

	// fetch order details
	order, err := h.service.GetOrderDetail(orderID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Order details retrieved successfully", order)
}

func (h *OrderHandler) GetUserTickets(c *gin.Context) {
	// extract ids
	orderID := c.Param("id")
	userID := utils.MustGetUserID(c)

	// fetch user tickets
	tickets, err := h.service.GetUserTicketsByOrder(orderID, userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "User tickets retrieved successfully", tickets)
}

func (h *OrderHandler) RefundOrder(c *gin.Context) {
	// extract user ID
	userID := utils.MustGetUserID(c)
	// extract order ID
	orderID := c.Param("id")

	// bind refund request
	var req dto.RefundOrderRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	// process order refund
	refundResult, err := h.service.RefundOrder(orderID, userID, req.Reason)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Order refunded successfully", refundResult)
}
