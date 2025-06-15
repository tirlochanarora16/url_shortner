package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", checkHealth)
	server.POST("/shorten", createShortUrl)
	server.PUT("/shorten/:shortCode", updateOriginalUrl)
	server.GET("/:shortCode", catchAllRoutes)
	server.DELETE("/shorten/:shortCode", deleteUrl)
}
