package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SimulationTrade struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Symbol     string    `gorm:"type:varchar(50);not null" json:"symbol"` // BTC/USD, AAPL, etc.
	TradeType  string    `gorm:"type:varchar(10);not null" json:"trade_type"` // buy, sell
	EntryPrice float64   `gorm:"type:decimal(18,8);not null" json:"entry_price"`
	Quantity   float64   `gorm:"type:decimal(18,8);not null" json:"quantity"`
	EntryAt    time.Time `gorm:"not null" json:"entry_at"`
	ExitPrice  *float64  `gorm:"type:decimal(18,8)" json:"exit_price"`
	ExitAt     *time.Time `json:"exit_at"`
	ProfitLoss *float64  `gorm:"type:decimal(18,8)" json:"profit_loss"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (s *SimulationTrade) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

type TradeDiary struct {
	ID                uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID            uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	TradeID           *uuid.UUID `gorm:"type:uuid" json:"trade_id"`
	GeneratedAnalysis string     `gorm:"type:text" json:"generated_analysis"`
	MarketContext     string     `gorm:"type:text" json:"market_context"`
	UserMemo          string     `gorm:"type:text" json:"user_memo"`
	DiaryDate         time.Time  `gorm:"type:date;not null" json:"diary_date"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (t *TradeDiary) BeforeCreate(tx *gorm.DB) error {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}
	return nil
}

type TradeRequest struct {
	Symbol     string  `json:"symbol" binding:"required"`
	TradeType  string  `json:"trade_type" binding:"required"` // buy, sell
	EntryPrice float64 `json:"entry_price" binding:"required"`
	Quantity   float64 `json:"quantity" binding:"required"`
}

type CloseTradeRequest struct {
	ExitPrice float64 `json:"exit_price" binding:"required"`
}
