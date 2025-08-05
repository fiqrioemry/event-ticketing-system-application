package repositories

import (
	"context"
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"gorm.io/gorm"
)

type AuditLogRepository interface {
	Create(ctx context.Context, log *models.AuditLog) error
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) Create(ctx context.Context, log *models.AuditLog) error {
	log.CreatedAt = time.Now()
	return r.db.WithContext(ctx).Create(log).Error
}
