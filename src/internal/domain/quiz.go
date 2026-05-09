package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizItem struct {
	ID            string    `gorm:"type:uuid;primaryKey" json:"id"`
	ConceptID     string    `gorm:"type:uuid;not null;index" json:"concept_id"`
	DifficultyIRT float64   `gorm:"not null" json:"difficulty_irt"`
	Discrimination float64  `gorm:"default:1.0" json:"discrimination"`
	Guessing      float64   `gorm:"default:0.25" json:"guessing"`
	Content       JSONMap   `gorm:"type:jsonb;not null" json:"content"`
	Modality      string    `gorm:"default:'text'" json:"modality"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Concept Concept `gorm:"foreignKey:ConceptID" json:"-"`
}

func (q *QuizItem) BeforeCreate(tx *gorm.DB) error {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	return nil
}

type Interaction struct {
	ID                  string    `gorm:"type:uuid;primaryKey" json:"id"`
	SessionID           string    `gorm:"type:uuid;not null;index" json:"session_id"`
	Type                string    `gorm:"not null" json:"type"`
	Payload             JSONMap   `gorm:"type:jsonb;not null" json:"payload"`
	WasCorrect          *bool     `json:"was_correct,omitempty"`
	ResponseTimeMs      *int      `json:"response_time_ms,omitempty"`
	RemediationGenerated *string  `gorm:"type:jsonb" json:"remediation_generated,omitempty"`
	CreatedAt           time.Time `json:"created_at"`

	Session Session `gorm:"foreignKey:SessionID" json:"-"`
}

func (i *Interaction) BeforeCreate(tx *gorm.DB) error {
	if i.ID == "" {
		i.ID = uuid.New().String()
	}
	return nil
}

type JSONMap map[string]interface{}

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan JSONMap: type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}
