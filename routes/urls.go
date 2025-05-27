package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tirlochanarora16/url_shortner/models"
)

func createShortUrl(c *gin.Context) {
	// body with receive original_url
	url, err := models.CheckUrlExists("https://example.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	// no result found for the given URL
	if url == nil {
		c.JSON(404, gin.H{"message": "No result found"})
		return
	}

	fmt.Println(url.ID)

	c.JSON(200, gin.H{"url": url})
}
