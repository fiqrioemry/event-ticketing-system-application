package handlers

import (
	"server/dto"
	"server/services"
	"server/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) CreateNewOrder(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.CreateOrderRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	orderID, err := h.service.CreateNewOrder(req, userID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"orderID": orderID})
}

func (h *OrderHandler) GetMyOrders(c *gin.Context) {

	userID := utils.MustGetUserID(c)

	var params dto.OrderQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	orders, pagination, err := h.service.GetMyOrders(userID, params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       orders,
		"pagination": pagination,
	})
}

func (h *OrderHandler) GetOrderDetail(c *gin.Context) {
	orderID := c.Param("id")

	order, err := h.service.GetOrderDetail(orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetUserTicketsByOrder(c *gin.Context) {
	orderID := c.Param("id")

	tickets, err := h.service.GetUserTicketsByOrderID(orderID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tickets,
	})
}

func (h *OrderHandler) RefundOrder(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	orderID := c.Param("id")

	var req dto.RefundOrderRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	res, err := h.service.RefundOrder(orderID, userID, req.Reason)
	if err != nil {
		utils.HandleServiceError(c, err, "refund failed")
		return
	}

	c.JSON(http.StatusOK, res)
}
