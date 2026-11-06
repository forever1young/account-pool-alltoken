package middleware

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RequireClientAPIKey guards the OpenAI-compatible /v1 routes against business clients.
func RequireClientAPIKey(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		provided := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if provided != "" && subtle.ConstantTimeCompare([]byte(provided), []byte(validKey)) == 1 {
			c.Next()
			return
		}
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
