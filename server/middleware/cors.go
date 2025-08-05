package middleware

import (
	"log"
	"slices"
	"strings"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		path := c.Request.URL.Path

		var allowedOrigin string

		if config.AppConfig.AppEnv == "production" {
			allowedOrigin = getProductionOrigin(origin)
			log.Printf("🔍 Production - Allowed Origin: '%s'", allowedOrigin)
			if origin != "" && allowedOrigin == "" {

				err := response.NewForbidden("Origin not allowed by CORS policy")
				response.Error(c, err)
				return
			}
		} else {
			allowedOrigin = getDevelopmentOrigin(origin)
		}

		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key, Accept, Origin, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			log.Printf("⚠️  OPTIONS request handled for: %s", path)
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func getProductionOrigin(origin string) string {
	if slices.Contains(config.AppConfig.AllowedOrigins, origin) {
		return origin
	}

	if len(config.AppConfig.AllowedOrigins) > 0 {
		return config.AppConfig.AllowedOrigins[0]
	}

	return ""
}

func getDevelopmentOrigin(origin string) string {
	if isLocalhost(origin) {
		return origin
	}

	if slices.Contains(config.AppConfig.AllowedOrigins, origin) {
		return origin
	}

	if len(config.AppConfig.AllowedOrigins) > 0 {
		return config.AppConfig.AllowedOrigins[0]
	}

	return "*"
}

func isLocalhost(origin string) bool {
	if origin == "" {
		return false
	}

	return strings.HasPrefix(origin, "http://localhost") ||
		strings.HasPrefix(origin, "https://localhost") ||
		strings.HasPrefix(origin, "http://127.0.0.1") ||
		strings.HasPrefix(origin, "https://127.0.0.1")
}
