package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"github.com/tirlochanarora16/url_shortner/models"
)

func createShortUrl(c *gin.Context) {
	// body with receive original_url
	var body models.NewShortUrlBody
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the original_url value",
			"error":   err.Error(),
		})
		return
	}

	url, err := models.CheckUrlExists(body.OriginalUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	// URL already exists in the DB
	if url != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Short URL for this URL already exists"})
		return
	}

	var newShortUrl models.Urls

	shortCode, err := shortid.Generate()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create short url for the given url",
			"error":   err.Error(),
		})
		return
	}

	newShortUrl.OriginalUrl = body.OriginalUrl
	newShortUrl.ShortCode = shortCode
	res, err := newShortUrl.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "unable to create short url for the given url",
			"error":   err.Error(),
		})
		return
	}

	// create New record in the DB
	c.JSON(http.StatusCreated, gin.H{"response": res})
}
