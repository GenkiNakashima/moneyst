package controllers

import (
	"net/http"

	"github.com/GenkiNakashima/moneyst/backend/config"
	"github.com/GenkiNakashima/moneyst/backend/middleware"
	"github.com/GenkiNakashima/moneyst/backend/models"
	"github.com/gin-gonic/gin"
)

func GetPersonalizedNews(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var news []models.SummarizedNews
	query := config.DB.Where("user_id = ? OR user_id IS NULL", userID).
		Order("published_at DESC").
		Limit(20)

	if err := query.Find(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
		return
	}

	c.JSON(http.StatusOK, news)
}

func MarkNewsAsRead(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	newsID := c.Param("id")
	var newsItem models.SummarizedNews

	if err := config.DB.Where("id = ?", newsID).First(&newsItem).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	// Ensure the news belongs to the user or is public
	if newsItem.UserID != nil && *newsItem.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	newsItem.ReadStatus = true
	if err := config.DB.Save(&newsItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update news"})
		return
	}

	c.JSON(http.StatusOK, newsItem)
}
