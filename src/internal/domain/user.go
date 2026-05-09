package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              string     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Email           string     `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash    string     `gorm:"not null" json:"-"`
	EstimatedTheta  *float64   `json:"estimated_theta,omitempty"`
	ThetaUncertainty *float64  `json:"theta_uncertainty,omitempty"`
	Cluster         *string    `gorm:"type:varchar(50)" json:"cluster,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
