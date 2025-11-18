package controllers

import (
	"net/http"

	"github.com/GenkiNakashima/moneyst/backend/config"
	"github.com/GenkiNakashima/moneyst/backend/middleware"
	"github.com/GenkiNakashima/moneyst/backend/models"
	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserProfile(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req models.UserProfileUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update fields if provided
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.RiskTolerance != "" {
		user.RiskTolerance = req.RiskTolerance
	}
	if req.LearningStyle != "" {
		user.LearningStyle = req.LearningStyle
	}
	if req.InvestmentExperience != "" {
		user.InvestmentExperience = req.InvestmentExperience
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}
