package domain

import (
	"time"
)

type Capsule struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Topic     string    `gorm:"not null" json:"topic"`
	Modalities string   `gorm:"type:jsonb;not null" json:"modalities"`
	A2UITree  string    `gorm:"type:jsonb;not null" json:"a2ui_tree"`
	SessionID *string   `gorm:"type:uuid" json:"session_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Session *Session `gorm:"foreignKey:SessionID" json:"-"`
}
