package repositories

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrderByID(ID string) (*models.Order, error)
	WithTx(fn func(tx *gorm.DB) (string, error)) (string, error)
	GetOrderDetails(orderID string) ([]models.OrderDetail, error)
	GetMyOrders(userID string, params dto.OrderQueryParams) ([]models.Order, int64, error)
	UpdateOrderStatus(orderID string, status string) error
	UpdateOrder(order *models.Order) error
	HasUsedTicket(orderID string) (bool, error)
	UpdatePaymentStatus(orderID string, status string) error
	IncreaseUserBalance(userID string, amount float64) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) WithTx(fn func(tx *gorm.DB) (string, error)) (string, error) {
	var result string
	err := r.db.Transaction(func(tx *gorm.DB) error {
		res, err := fn(tx)
		result = res
		return err
	})
	return result, err
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetOrderByID(ID string) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Event").First(&order, "id = ?", ID).Error
	return &order, err
}

func (r *orderRepository) GetOrderDetails(orderID string) ([]models.OrderDetail, error) {
	var orderDetails []models.OrderDetail
	err := r.db.Where("order_id = ?", orderID).Find(&orderDetails).Error
	return orderDetails, err
}

func (r *orderRepository) GetMyOrders(userID string, params dto.OrderQueryParams) ([]models.Order, int64, error) {
	var orders []models.Order
	var count int64

	// Base query dengan JOIN explicit
	db := r.db.Model(&models.Order{}).
		Joins("LEFT JOIN events ON orders.event_id = events.id").
		Where("orders.user_id = ?", userID)

	// Search by event name atau description
	if params.Q != "" {
		like := "%" + params.Q + "%"
		db = db.Where("events.title LIKE ? OR events.description LIKE ?", like, like)
	}

	// Status filter
	if params.Status != "" && params.Status != "all" {
		db = db.Where("orders.status = ?", params.Status)
	}

	// Sorting dengan table prefix
	switch params.Sort {
	case "name_asc":
		db = db.Order("events.title ASC")
	case "name_desc":
		db = db.Order("events.title DESC")
	case "price_asc":
		db = db.Order("orders.total_price ASC")
	case "price_desc":
		db = db.Order("orders.total_price DESC")
	case "created_at_asc":
		db = db.Order("orders.created_at ASC")
	case "created_at_desc":
		db = db.Order("orders.created_at DESC")
	default:
		db = db.Order("orders.created_at DESC")
	}

	// Count total dengan field yang tepat
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Pagination dan fetch result dengan Preload
	offset := (params.Page - 1) * params.Limit
	if err := db.Preload("Event").
		Limit(params.Limit).
		Offset(offset).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, count, nil
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) UpdatePaymentStatus(orderID string, status string) error {
	return r.db.Model(&models.Payment{}).Where("order_id = ?", orderID).Update("status", status).Error
}

func (r *orderRepository) IncreaseUserBalance(userID string, amount float64) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func (r *orderRepository) HasUsedTicket(orderID string) (bool, error) {
	var count int64
	err := r.db.Table("user_tickets").
		Joins("JOIN order_details ON user_tickets.ticket_id = order_details.ticket_id").
		Where("order_details.order_id = ? AND user_tickets.is_used = ?", orderID, true).
		Count(&count).Error
	return count > 0, err
}

func (r *orderRepository) UpdateOrderStatus(orderID string, status string) error {
	return r.db.Model(&models.Order{}).
		Where("id = ?", orderID).
		Update("status", status).Error
}
