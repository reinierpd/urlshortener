package main

import (
	"UrlMinifier/shortener"
	"UrlMinifier/store"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

func createShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.AddUrl(shortUrl, creationRequest.LongUrl)

	// TODO: store host in a config file
	host := "http://localhost:8080/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}
func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	longlUrl := store.GetLongUrl(shortUrl)
	c.Redirect(302, longlUrl)
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
	r.POST("/create-short-url", func(c *gin.Context) {
		createShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		HandleShortUrlRedirect(c)
	})

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
