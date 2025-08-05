package utils

import (
	"encoding/json"
	"fmt"

	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func BuildAuditLog(c *gin.Context, userID string, action string, resource string, payload any) *models.AuditLog {
	description := buildDescription(action, resource, payload)

	return &models.AuditLog{
		ID:          uuid.NewString(),
		UserID:      userID,
		Action:      action,
		Resource:    resource,
		Description: description,
		IP:          c.ClientIP(),
		UserAgent:   c.Request.UserAgent(),
	}
}

func buildDescription(action string, resource string, payload any) string {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Sprintf("User performed %s on %s", action, resource)
	}
	return fmt.Sprintf("User performed %s on %s with data: %s", action, resource, string(jsonData))
}
