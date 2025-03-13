package entity

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageID   uuid.UUID `json:"message_id" gorm:"type:uuid;primaryKey"`
	Message     string    `json:"message" gorm:"type:text;not null"`
	CommunityId uuid.UUID `json:"community_id" gorm:"type:uuid;not null"`
	UserId      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	TimeSent    time.Time `json:"time_sent" gorm:"type:timestamp;default:current_timestamp"`
	User        User      `gorm:"foreignKey:UserId"`
	Community   Community `gorm:"foreignKey:CommunityID"`
}
