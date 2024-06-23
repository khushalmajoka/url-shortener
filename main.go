package main

import (
	"fmt"
	"url-shortener/config"
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/shorten", handlers.ShortenURL)
	router.GET("/:shortURL", handlers.RedirectURL)
	router.GET("/metrics", handlers.GetMetrics)

	router.Run(fmt.Sprintf(":%s", config.Port))
}
