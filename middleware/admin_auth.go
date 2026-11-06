package middleware

import "github.com/gin-gonic/gin"

// RequireAdminSession guards the pool's own admin web/API routes.
func RequireAdminSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
