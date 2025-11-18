package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email                string    `gorm:"uniqueIndex;not null" json:"email"`
	HashedPassword       string    `gorm:"not null" json:"-"`
	Username             string    `gorm:"not null" json:"username"`
	RiskTolerance        string    `gorm:"type:varchar(50)" json:"risk_tolerance"`        // low, medium, high
	LearningStyle        string    `gorm:"type:varchar(50)" json:"learning_style"`        // text, visual, conversational
	InvestmentExperience string    `gorm:"type:varchar(50)" json:"investment_experience"` // none, beginner, intermediate
	TradingPatternMemo   string    `gorm:"type:text" json:"trading_pattern_memo"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserProfileUpdate struct {
	Username             string `json:"username"`
	RiskTolerance        string `json:"risk_tolerance"`
	LearningStyle        string `json:"learning_style"`
	InvestmentExperience string `json:"investment_experience"`
}
