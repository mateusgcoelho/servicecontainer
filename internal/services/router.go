package services

import "github.com/gin-gonic/gin"

func SetupRoutes(ctx *gin.Engine) {
	r := ctx.Group("/services")

	r.GET("/", handleGetServices)
	r.POST("/:tag/active", handleActiveService)
}
