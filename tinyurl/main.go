package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Drime648/coding-challenges/tinyurl/internal/storage"
	"github.com/Drime648/coding-challenges/tinyurl/handler"
)


func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/generate", func (c *gin.Context) {
		handler.GenerateUrl(c)
	})

	r.GET("/:shortUrl", func (c *gin.Context) {
		handler.HandleRedirect(c)
	})

	err := storage.InitStorage()
	if err != nil {
		panic(err)
	}

	err = r.Run(":9090")
	if err != nil {
		panic(fmt.Sprintf("Could not start web server: %v\n", err))
	}
}
