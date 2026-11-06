package controller

import "github.com/gin-gonic/gin"

// ChatCompletions implements the OpenAI-compatible /v1/chat/completions endpoint.
func ChatCompletions(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}

// ListModels implements the OpenAI-compatible /v1/models endpoint.
func ListModels(c *gin.Context) {
	c.JSON(501, gin.H{"message": "not implemented"})
}
