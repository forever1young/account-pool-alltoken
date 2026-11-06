package controller

import "github.com/gin-gonic/gin"

// ListAccounts returns pool account status for the admin UI.
func ListAccounts(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

// ListUsage returns usage/billing summaries for the admin UI.
func ListUsage(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}
