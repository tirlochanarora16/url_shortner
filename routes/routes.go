package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "running successfully!"})
	})
	server.POST("/shorten", createShortUrl)
	// server.GET("/*shortCode", catchAllRoutes)
}
