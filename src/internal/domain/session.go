package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Session struct {
	ID                string         `gorm:"type:uuid;primaryKey" json:"id"`
	UserID            string         `gorm:"type:uuid;not null;index" json:"user_id"`
	State             string         `gorm:"not null;default:'coldstart'" json:"state"`
	A2UISnapshot      *string        `gorm:"type:jsonb" json:"a2ui_snapshot,omitempty"`
	TargetSuccessRate float64        `gorm:"not null;default:0.85" json:"target_success_rate"`
	StartedAt         time.Time      `gorm:"not null" json:"started_at"`
	CompletedAt       *time.Time     `json:"completed_at,omitempty"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	User User `gorm:"foreignKey:UserID" json:"-"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}
