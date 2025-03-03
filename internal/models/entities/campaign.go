package entity

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	CampaignId      uuid.UUID `json:"campaign_id" gorm:"type:uuid;primaryKey"`
	Title           string    `json:"title" gorm:"type:varchar(255)"`
	Description     string    `json:"description" gorm:"type:text"`
	GoalAmount      int       `json:"goal_amount" gorm:"type:integer;check:goal_amount>0"`
	CollectedAmount int       `json:"collected_amount" gorm:"type:integer;default:0"`
	IsActive        bool      `json:"is_active" gorm:"type:bool;default:true"`
	CreatedAt       time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
	UserId          uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User            User      `gorm:"foreignKey:UserId"`
}
