package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Concept struct {
	ID          string     `gorm:"type:uuid;primaryKey" json:"id"`
	ParentID    *string    `gorm:"type:uuid" json:"parent_id,omitempty"`
	Name        string     `gorm:"not null" json:"name"`
	Description *string    `json:"description,omitempty"`
	Difficulty  float64    `gorm:"not null;default:0.5" json:"difficulty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	Parent    *Concept   `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children  []Concept  `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	QuizItems []QuizItem `gorm:"foreignKey:ConceptID" json:"quiz_items,omitempty"`
}

func (c *Concept) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

type UserConceptMastery struct {
	UserID         string     `gorm:"type:uuid;primaryKey" json:"user_id"`
	ConceptID      string     `gorm:"type:uuid;primaryKey" json:"concept_id"`
	PLearned       float64    `gorm:"not null;default:0.3" json:"p_learned"`
	PGuess         float64    `gorm:"not null;default:0.1" json:"p_guess"`
	PSlip          float64    `gorm:"not null;default:0.1" json:"p_slip"`
	PTransit       float64    `gorm:"not null;default:0.2" json:"p_transit"`
	LastPracticed  *time.Time `json:"last_practiced,omitempty"`
	EasinessFactor float64    `gorm:"default:2.5" json:"easiness_factor"`
	IntervalDays   int        `gorm:"default:1" json:"interval_days"`
	Repetitions    int        `gorm:"default:0" json:"repetitions"`
	NextReview     *time.Time `json:"next_review,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`

	User    User    `gorm:"foreignKey:UserID" json:"-"`
	Concept Concept `gorm:"foreignKey:ConceptID" json:"concept,omitempty"`
}
