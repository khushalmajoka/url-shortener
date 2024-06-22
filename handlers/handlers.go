package handlers

import (
	"net/http"
	"url-shortener/storage"
	"url-shortener/utils"
	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	var request struct {
		URL string `json:"url"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortURL := utils.GenerateShortURL(request.URL)
	storage.SaveURL(request.URL, shortURL)
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	originalURL := storage.GetOriginalURL(shortURL)
	if originalURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func GetMetrics(c *gin.Context) {
	metrics := storage.GetTopDomains(3)
	c.JSON(http.StatusOK, metrics)
}
