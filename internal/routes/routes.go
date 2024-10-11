package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/health", Health)
	router.GET("/ready", Ready)

	api := router.Group("/agriculture")
	api.GET("/summary", Agriculture)
}
