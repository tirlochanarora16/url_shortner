package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"github.com/tirlochanarora16/url_shortner/models"
	"github.com/tirlochanarora16/url_shortner/pkg"
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

	isValidUrl := models.IsValidUrl(body.OriginalUrl)

	if !isValidUrl {
		log.Println("Invalid URL provided")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "the url provided is invalid",
		})
		return
	}

	url, err := models.CheckUrlExists(body.OriginalUrl)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Short URL for this URL already exists"})
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

func catchAllRoutes(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := models.CheckShortCode(shortCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid url",
			"error":   err.Error(),
		})
		return
	}

	go func() {
		key := fmt.Sprintf("url_hits:%s", shortCode)
		err := pkg.Rdb.Incr(pkg.Ctx, key).Err()

		if err != nil {
			log.Printf("Redis increment failed for %s: %v", shortCode, err)
		}
	}()

	c.Redirect(http.StatusFound, url.OriginalUrl)
}

func updateOriginalUrl(c *gin.Context) {
	var requestBody models.NewShortUrlBody
	c.ShouldBindJSON(&requestBody)
	shortCode := c.Param("shortCode")
	url, err := models.CheckShortCode(shortCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid url",
			"error":   err.Error(),
		})
		return
	}

	isValidUrl := models.IsValidUrl(requestBody.OriginalUrl)

	if !isValidUrl {
		log.Println("Invalid URL provided")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "the url provided is invalid",
		})
		return
	}

	res, err := url.UpdateUrl(map[string]interface{}{
		"original_url": requestBody.OriginalUrl,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"res": res,
	})
}

func deleteUrl(c *gin.Context) {
	shortCode := c.Param("shortCode")
	url, err := models.CheckShortCode(shortCode)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "invalid short url given",
		})
		return
	}

	err = url.Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "deleted successfully",
	})
}
