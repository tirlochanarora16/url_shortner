package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", checkHealth)
	server.POST("/shorten", createShortUrl)
	server.PUT("/shorten/:shortcode", updateOriginalUrl)
	server.GET("/:shortcode", catchAllRoutes)
	server.DELETE("/shorten/:shortcode", deleteUrl)
}
