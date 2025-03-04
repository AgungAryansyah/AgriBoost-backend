package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id            uuid.UUID     `json:"user_id" gorm:"type:uuid;primaryKey"`
	Name          string        `json:"name" gorm:"type:varchar(255)"`
	Email         string        `json:"email" gorm:"type:varchar(255);unique"`
	Password      string        `json:"password" gorm:"type:varchar(255)"`
	QuizPoint     int           `json:"quiz_point" gorm:"type:integer;default:0"`
	DonationPoint int           `json:"donation_point" gorm:"type:integer;default:0"`
	IsAdmin       bool          `json:"is_admin" gorm:"type:bool;default:false"`
	CreatedAt     time.Time     `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time     `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
	QuizAttempt   []QuizAttempt `gorm:"foreignKey:Id"`
	Donation      []Donation    `gorm:"foreignKey:Id"`
	Campaign      []Campaign    `gorm:"foreignKey:Id"`
}
