package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", checkHealth)
	server.POST("/shorten", createShortUrl)
	server.GET("/:shortcode", catchAllRoutes)
}
