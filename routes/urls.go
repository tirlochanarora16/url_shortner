package routes

import "github.com/gin-gonic/gin"

func createShortUrl(c *gin.Context) {
	c.JSON(500, gin.H{"m": "hello world 123"})
}
