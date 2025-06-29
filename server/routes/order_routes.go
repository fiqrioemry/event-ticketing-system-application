package routes

import (
	"server/handlers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup, h *handlers.OrderHandler) {
	order := r.Group("/orders", middleware.AuthRequired(), middleware.RoleOnly("user"))

	order.GET("", h.GetMyOrders)
	order.GET("/:id", h.GetOrderDetail)
	order.POST("", h.CreateNewOrder)
	order.GET("/:id/user-tickets", h.GetUserTickets)
	order.POST("/:id/refund", h.RefundOrder)

}
