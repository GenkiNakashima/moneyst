package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Curriculum struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title      string    `gorm:"not null" json:"title"`
	ModuleType string    `gorm:"type:varchar(100)" json:"module_type"` // chart, news, simulation, quiz
	ContentID  uuid.UUID `gorm:"type:uuid" json:"content_id"`
	Status     string    `gorm:"type:varchar(50);default:'pending'" json:"status"` // pending, in_progress, completed
	Order      int       `gorm:"not null" json:"order"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (c *Curriculum) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type Lesson struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title             string    `gorm:"not null" json:"title"`
	ContentBody       string    `gorm:"type:text;not null" json:"content_body"`
	DifficultyLevel   int       `json:"difficulty_level"`                                   // 1-5
	Category          string    `gorm:"type:varchar(100)" json:"category"`                  // chart, glossary, economy, quiz
	EstimatedTimeMin  int       `json:"estimated_time_min"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (l *Lesson) BeforeCreate(tx *gorm.DB) error {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return nil
}

type DailyLesson struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	LessonDate   time.Time `gorm:"type:date;not null" json:"lesson_date"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	Completed    bool      `gorm:"default:false" json:"completed"`
	PointsEarned int       `gorm:"default:0" json:"points_earned"`
	CreatedAt    time.Time `json:"created_at"`
}

func (d *DailyLesson) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}
