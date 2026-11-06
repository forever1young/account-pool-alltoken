package router

import (
	"github.com/gin-gonic/gin"

	"github.com/forever1young/account-pool-alltoken/config"
	"github.com/forever1young/account-pool-alltoken/controller"
	"github.com/forever1young/account-pool-alltoken/middleware"
)

func New(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1", middleware.RequireClientAPIKey(cfg.ClientAPIKey))
	{
		v1.POST("/chat/completions", controller.ChatCompletions)
		v1.GET("/models", controller.ListModels)
	}

	admin := r.Group("/api/admin", middleware.RequireAdminSession())
	{
		admin.GET("/accounts", controller.ListAccounts)
		admin.GET("/usage", controller.ListUsage)
	}

	return r
}
