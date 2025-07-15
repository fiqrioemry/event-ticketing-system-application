package middleware

import (
	"fmt"
	"slices"

	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("accessToken")
		if err != nil || tokenString == "" {
			response.Error(c, response.Unauthorized("Unauthorized!! Token missing"))
			c.Abort()
			return
		}

		claims, err := utils.DecodeAccessToken(tokenString)
		if err != nil {
			response.Error(c, response.Unauthorized("Invalid or expired token"))
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Set("userID", claims.UserID)

		c.Next()
	}
}

func RoleOnly(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := utils.MustGetRole(c)
		if slices.Contains(allowedRoles, role) {
			c.Next()
			return
		}
		response.Error(c, response.Forbidden(fmt.Sprintf("Forbidden: Access denied. Only roles %v are allowed", allowedRoles)))
		c.Abort()
	}
}
