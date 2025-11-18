package controllers

import (
	"net/http"
	"time"

	"github.com/GenkiNakashima/moneyst/backend/config"
	"github.com/GenkiNakashima/moneyst/backend/middleware"
	"github.com/GenkiNakashima/moneyst/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMyCurriculum(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var curriculum []models.Curriculum
	if err := config.DB.Where("user_id = ?", userID).Order("\"order\" ASC").Find(&curriculum).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch curriculum"})
		return
	}

	// If curriculum is empty, generate initial curriculum
	if len(curriculum) == 0 {
		curriculum = generateInitialCurriculum(userID)
		for _, item := range curriculum {
			config.DB.Create(&item)
		}
	}

	c.JSON(http.StatusOK, curriculum)
}

func CompleteCurriculumItem(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	itemID := c.Param("id")
	var item models.Curriculum

	if err := config.DB.Where("id = ? AND user_id = ?", itemID, userID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Curriculum item not found"})
		return
	}

	item.Status = "completed"
	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update curriculum"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func GetDailyLesson(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	today := time.Now().Format("2006-01-02")
	var dailyLesson models.DailyLesson

	err = config.DB.Where("user_id = ? AND lesson_date = ?", userID, today).First(&dailyLesson).Error
	if err != nil {
		// Generate new daily lesson
		dailyLesson = models.DailyLesson{
			UserID:     userID,
			LessonDate: time.Now(),
			Content:    generateDailyLessonContent(),
			Completed:  false,
		}
		config.DB.Create(&dailyLesson)
	}

	c.JSON(http.StatusOK, dailyLesson)
}

func CompleteDailyLesson(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	lessonID := c.Param("id")
	var lesson models.DailyLesson

	if err := config.DB.Where("id = ? AND user_id = ?", lessonID, userID).First(&lesson).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Daily lesson not found"})
		return
	}

	lesson.Completed = true
	lesson.PointsEarned = 10 // Award 10 points for completion
	if err := config.DB.Save(&lesson).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update lesson"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// Helper function to generate initial curriculum
func generateInitialCurriculum(userID uuid.UUID) []models.Curriculum {
	return []models.Curriculum{
		{UserID: userID, Title: "Introduction to Investment", ModuleType: "quiz", Status: "pending", Order: 1},
		{UserID: userID, Title: "Understanding Charts", ModuleType: "chart", Status: "pending", Order: 2},
		{UserID: userID, Title: "Market News Analysis", ModuleType: "news", Status: "pending", Order: 3},
		{UserID: userID, Title: "First Simulation Trade", ModuleType: "simulation", Status: "pending", Order: 4},
	}
}

// Helper function to generate daily lesson content
func generateDailyLessonContent() string {
	// TODO: Integrate with AI service
	return `# Today's Lesson: Understanding Support and Resistance

Support and resistance are key concepts in technical analysis:

- **Support**: A price level where buying pressure is strong enough to prevent the price from falling further.
- **Resistance**: A price level where selling pressure is strong enough to prevent the price from rising further.

**Quick Quiz**: Look at today's chart and identify one support level and one resistance level.`
}
