package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Post represents a user-generated post in the community
type Post struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID            uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User              *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Content           string    `gorm:"type:text;not null" json:"content"`
	IsFactChecked     bool      `gorm:"default:false" json:"is_fact_checked"`
	FactCheckResult   string    `gorm:"type:text" json:"fact_check_result,omitempty"`
	FactCheckStatus   string    `gorm:"type:varchar(50);default:'pending'" json:"fact_check_status"` // pending, approved, flagged, removed
	LikesCount        int       `gorm:"-" json:"likes_count"`                                         // Computed field
	RepliesCount      int       `gorm:"-" json:"replies_count"`                                       // Computed field
	UserLiked         bool      `gorm:"-" json:"user_liked"`                                          // Computed field for current user
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// Reply represents a response to a post
type Reply struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	PostID       uuid.UUID `gorm:"type:uuid;not null" json:"post_id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User         *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	IsAIResponse bool      `gorm:"default:false" json:"is_ai_response"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// PostLike represents a like/upvote on a post
type PostLike struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	PostID    uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_post_user_like" json:"post_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_post_user_like" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// PostWithDetails includes post with related data
type PostWithDetails struct {
	Post
	Replies []Reply `json:"replies,omitempty"`
}

// TrendingPost represents a post in the trending list
type TrendingPost struct {
	Post
	Rank int `json:"rank"`
}

// BeforeCreate hooks for UUID generation
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	if p.FactCheckStatus == "" {
		p.FactCheckStatus = "pending"
	}
	return nil
}

func (r *Reply) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (pl *PostLike) BeforeCreate(tx *gorm.DB) error {
	if pl.ID == uuid.Nil {
		pl.ID = uuid.New()
	}
	return nil
}

// Request/Response structs
type CreatePostRequest struct {
	Content string `json:"content" binding:"required,min=1,max=5000"`
}

type CreateReplyRequest struct {
	Content string `json:"content" binding:"required,min=1,max=2000"`
}

type SearchPostsRequest struct {
	Query  string `json:"query" form:"query" binding:"required,min=1"`
	Limit  int    `json:"limit" form:"limit"`
	Offset int    `json:"offset" form:"offset"`
}
