package routes

import (
	"github.com/GenkiNakashima/moneyst/backend/controllers"
	"github.com/GenkiNakashima/moneyst/backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// User profile
		api.GET("/auth/me", controllers.GetMe)
		api.GET("/users/me/profile", controllers.GetUserProfile)
		api.PUT("/users/me/profile", controllers.UpdateUserProfile)

		// Curriculum
		api.GET("/curriculums/me", controllers.GetMyCurriculum)
		api.POST("/curriculums/:id/complete", controllers.CompleteCurriculumItem)

		// Daily lessons
		api.GET("/lessons/daily", controllers.GetDailyLesson)
		api.POST("/lessons/:id/complete", controllers.CompleteDailyLesson)

		// Simulation trades
		api.POST("/simulations/trades", controllers.CreateSimulationTrade)
		api.GET("/simulations/trades", controllers.GetSimulationTrades)
		api.POST("/simulations/trades/:id/close", controllers.CloseSimulationTrade)

		// Trade diaries
		api.GET("/simulations/diaries", controllers.GetTradeDiaries)
		api.POST("/simulations/trades/:id/diary", controllers.GenerateTradeDiary)
		api.PUT("/simulations/diaries/:id", controllers.UpdateTradeDiaryMemo)

		// News
		api.GET("/news/personalized", controllers.GetPersonalizedNews)
		api.POST("/news/:id/read", controllers.MarkNewsAsRead)

		// Community
		api.POST("/community/posts", controllers.CreatePost)
		api.GET("/community/posts", controllers.GetPosts)
		api.GET("/community/posts/trending", controllers.GetTrendingPosts)
		api.GET("/community/posts/search", controllers.SearchPosts)
		api.GET("/community/posts/:id", controllers.GetPost)
		api.DELETE("/community/posts/:id", controllers.DeletePost)
		api.POST("/community/posts/:id/replies", controllers.CreateReply)
		api.POST("/community/posts/:id/like", controllers.ToggleLike)
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
