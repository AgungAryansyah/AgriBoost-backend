package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id              uuid.UUID         `json:"user_id" gorm:"type:uuid;primaryKey"`
	Name            string            `json:"name" gorm:"type:varchar(255);notnull;unique"`
	ProfilePicture  string            `json:"profile_picture" gorm:"type:text;unique"`
	Phone           string            `json:"phone" gorm:"type:varchar(15);notnull;unique"`
	Email           string            `json:"email" gorm:"type:varchar(255);unique"`
	Password        string            `json:"password" gorm:"type:varchar(255);notnull"`
	QuizPoint       int               `json:"quiz_point" gorm:"type:integer;notnull;default:0"`
	DonationPoint   int               `json:"donation_point" gorm:"type:integer;notnull;default:0"`
	IsAdmin         bool              `json:"is_admin" gorm:"type:bool;notnull;default:false"`
	CreatedAt       time.Time         `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt       time.Time         `json:"updated_at" gorm:"type:timestamp;default:current_timestamp"`
	Article         Article           `gorm:"foreignKey:UserId"`
	QuizAttempt     []QuizAttempt     `gorm:"foreignKey:UserId"`
	Donation        []Donation        `gorm:"foreignKey:UserId"`
	Campaign        []Campaign        `gorm:"foreignKey:UserId"`
	CommunityMember []CommunityMember `gorm:"foreignKey:UserId"`
}
