package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

func LimitFileSize(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)
		c.Next()
	}
}

func RateLimiter(maxAttempts int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := GetClientIP(c)
		key := fmt.Sprintf("ratelimit:%s", ip)

		count, _ := config.RedisClient.Get(config.Ctx, key).Int()
		if count >= maxAttempts {
			parsedErr := response.NewTooManyRequests("Too many requests. Slow down baby.").
				WithContext("ip_address", ip)
			response.Error(c, parsedErr)
			c.Abort()
			return
		}

		pipe := config.RedisClient.TxPipeline()
		pipe.Incr(config.Ctx, key)
		pipe.Expire(config.Ctx, key, duration)
		_, _ = pipe.Exec(config.Ctx)

		c.Next()
	}
}

func GetClientIP(c *gin.Context) string {
	ip := c.ClientIP()
	if ip == "" {
		ip = c.Request.Header.Get("X-Forwarded-For")
		ip = strings.TrimSpace(strings.Split(ip, ",")[0])
	}
	return ip
}
