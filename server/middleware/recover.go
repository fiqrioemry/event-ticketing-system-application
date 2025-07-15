package middleware

import (
	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		response.Error(c, response.InternalServerError("An unexpected error occurred", nil))
		c.Abort()
	})
}
