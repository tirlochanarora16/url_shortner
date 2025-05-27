package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tirlochanarora16/url_shortner/models"
)

func createShortUrl(c *gin.Context) {
	// body with receive original_url
	url, err := models.CheckUrlExists("https://examples.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	// URL already exists in the DB
	if url != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Short URL for this URL already exists"})
		return
	}

	// create New record in the DB
	c.JSON(200, gin.H{"url": url})
}
