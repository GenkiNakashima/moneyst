package controllers

import (
	"net/http"
	"strconv"

	"github.com/GenkiNakashima/moneyst/backend/config"
	"github.com/GenkiNakashima/moneyst/backend/models"
	"github.com/GenkiNakashima/moneyst/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var aiService = services.NewAIService()

// CreatePost creates a new post with AI fact-checking
func CreatePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req models.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perform AI fact-checking
	factCheckResult, status, err := aiService.FactCheckContent(req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fact-check content"})
		return
	}

	// If content is flagged for removal, don't create the post
	if status == "removed" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":             "投稿が拒否されました",
			"fact_check_result": factCheckResult,
		})
		return
	}

	post := models.Post{
		UserID:          userID.(uuid.UUID),
		Content:         req.Content,
		IsFactChecked:   true,
		FactCheckResult: factCheckResult,
		FactCheckStatus: status,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// Load user info
	config.DB.Preload("User").First(&post, post.ID)

	// Check if @checkAI was invoked
	if aiService.CheckIfAIInvoked(req.Content) {
		// Generate AI response as a reply
		aiResponse, err := aiService.GenerateAIResponse(req.Content)
		if err == nil {
			reply := models.Reply{
				PostID:       post.ID,
				UserID:       userID.(uuid.UUID), // System user ID could be used here
				Content:      aiResponse,
				IsAIResponse: true,
			}
			config.DB.Create(&reply)
		}
	}

	c.JSON(http.StatusCreated, post)
}

// GetPosts retrieves paginated posts
func GetPosts(c *gin.Context) {
	userID, _ := c.Get("user_id")

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []models.Post
	query := config.DB.Preload("User").
		Where("fact_check_status != ?", "removed").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset)

	if err := query.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	// Enrich posts with counts and user-specific data
	for i := range posts {
		enrichPostData(&posts[i], userID)
	}

	c.JSON(http.StatusOK, posts)
}

// GetPost retrieves a single post with replies
func GetPost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	postID := c.Param("id")

	var post models.Post
	if err := config.DB.Preload("User").First(&post, "id = ?", postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Get replies
	var replies []models.Reply
	config.DB.Preload("User").
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Find(&replies)

	enrichPostData(&post, userID)

	c.JSON(http.StatusOK, gin.H{
		"post":    post,
		"replies": replies,
	})
}

// CreateReply creates a reply to a post
func CreateReply(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	postID := c.Param("id")

	var req models.CreateReplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify post exists
	var post models.Post
	if err := config.DB.First(&post, "id = ?", postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	reply := models.Reply{
		PostID:  post.ID,
		UserID:  userID.(uuid.UUID),
		Content: req.Content,
	}

	if err := config.DB.Create(&reply).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reply"})
		return
	}

	// Load user info
	config.DB.Preload("User").First(&reply, reply.ID)

	// Check if @checkAI was invoked
	if aiService.CheckIfAIInvoked(req.Content) {
		// Generate AI response
		aiResponse, err := aiService.GenerateAIResponse(req.Content)
		if err == nil {
			aiReply := models.Reply{
				PostID:       post.ID,
				UserID:       userID.(uuid.UUID),
				Content:      aiResponse,
				IsAIResponse: true,
			}
			config.DB.Create(&aiReply)
		}
	}

	c.JSON(http.StatusCreated, reply)
}

// ToggleLike toggles a like on a post
func ToggleLike(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	postID := c.Param("id")

	// Check if like exists
	var like models.PostLike
	err := config.DB.Where("post_id = ? AND user_id = ?", postID, userID).First(&like).Error

	if err == nil {
		// Unlike: delete the like
		config.DB.Delete(&like)
		c.JSON(http.StatusOK, gin.H{"liked": false, "message": "Like removed"})
	} else {
		// Like: create new like
		newLike := models.PostLike{
			PostID: uuid.MustParse(postID),
			UserID: userID.(uuid.UUID),
		}
		if err := config.DB.Create(&newLike).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like post"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"liked": true, "message": "Post liked"})
	}
}

// SearchPosts searches posts by content
func SearchPosts(c *gin.Context) {
	userID, _ := c.Get("user_id")
	query := c.Query("q")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	var posts []models.Post
	err := config.DB.Preload("User").
		Where("content ILIKE ? AND fact_check_status != ?", "%"+query+"%", "removed").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&posts).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search posts"})
		return
	}

	// Enrich posts with counts
	for i := range posts {
		enrichPostData(&posts[i], userID)
	}

	c.JSON(http.StatusOK, posts)
}

// GetTrendingPosts gets top 3 posts by likes count
func GetTrendingPosts(c *gin.Context) {
	userID, _ := c.Get("user_id")

	type PostWithLikeCount struct {
		models.Post
		LikeCount int `json:"like_count"`
	}

	var posts []PostWithLikeCount
	err := config.DB.Table("posts").
		Select("posts.*, COUNT(post_likes.id) as like_count").
		Joins("LEFT JOIN post_likes ON post_likes.post_id = posts.id").
		Where("posts.fact_check_status != ?", "removed").
		Group("posts.id").
		Order("like_count DESC, posts.created_at DESC").
		Limit(3).
		Find(&posts).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch trending posts"})
		return
	}

	// Load user info and enrich data
	for i := range posts {
		config.DB.Preload("User").First(&posts[i].Post, posts[i].ID)
		enrichPostData(&posts[i].Post, userID)
	}

	c.JSON(http.StatusOK, posts)
}

// enrichPostData adds computed fields to a post
func enrichPostData(post *models.Post, userID interface{}) {
	// Count likes
	var likeCount int64
	config.DB.Model(&models.PostLike{}).Where("post_id = ?", post.ID).Count(&likeCount)
	post.LikesCount = int(likeCount)

	// Count replies
	var replyCount int64
	config.DB.Model(&models.Reply{}).Where("post_id = ?", post.ID).Count(&replyCount)
	post.RepliesCount = int(replyCount)

	// Check if current user liked this post
	if userID != nil {
		var like models.PostLike
		err := config.DB.Where("post_id = ? AND user_id = ?", post.ID, userID).First(&like).Error
		post.UserLiked = (err == nil)
	}
}

// DeletePost deletes a user's own post
func DeletePost(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	postID := c.Param("id")

	var post models.Post
	if err := config.DB.First(&post, "id = ?", postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Check if user owns the post
	if post.UserID != userID.(uuid.UUID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own posts"})
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
