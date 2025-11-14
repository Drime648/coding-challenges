package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/Drime648/coding-challenges/tinyurl/internal/generator"
	"github.com/Drime648/coding-challenges/tinyurl/internal/storage"
	"net/http"
	"fmt"
)

var HOST = "http://localhost:9090/"

type GenerateRequest struct {
	Url string `json:"url" binding:"required"`
}

func GenerateUrl(c *gin.Context) {

	genRequest := &GenerateRequest{}

	err := c.ShouldBindJSON(genRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl, err := generator.GenerateUrl(genRequest.Url)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL"})
		return
	}

	err = storage.SaveUrl(genRequest.Url, shortUrl)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "shortened URL generated successfully",
		"short_url": HOST + shortUrl,
	})

}

func HandleRedirect(c *gin.Context) {

	shortUrl := c.Param("shortUrl")
	originalUrl, err := storage.GetUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid short URL"})
	}
	c.Redirect(http.StatusFound, originalUrl)
}
