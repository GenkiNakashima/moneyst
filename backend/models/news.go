package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SummarizedNews struct {
	ID                uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID            *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	OriginalURL       string     `gorm:"type:varchar(1000)" json:"original_url"`
	OriginalTitle     string     `gorm:"type:varchar(500)" json:"original_title"`
	SummarizedTitle   string     `gorm:"type:varchar(500)" json:"summarized_title"`
	SummarizedContent string     `gorm:"type:text" json:"summarized_content"`
	AIInsight         string     `gorm:"type:text" json:"ai_insight"`
	PublishedAt       time.Time  `gorm:"not null" json:"published_at"`
	ReadStatus        bool       `gorm:"default:false" json:"read_status"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (n *SummarizedNews) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return nil
}
