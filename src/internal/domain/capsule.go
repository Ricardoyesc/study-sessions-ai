package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Capsule struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Topic     string    `gorm:"not null" json:"topic"`
	Modalities string   `gorm:"type:jsonb;not null" json:"modalities"`
	A2UITree  string    `gorm:"type:jsonb;not null" json:"a2ui_tree"`
	SessionID *string   `gorm:"type:uuid" json:"session_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Session *Session `gorm:"foreignKey:SessionID" json:"-"`
}

func (c *Capsule) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}
