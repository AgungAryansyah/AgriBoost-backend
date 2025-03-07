package entity

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ArticleId    uuid.UUID `json:"article_id" gorm:"type:uuid;primaryKey"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	Title        string    `json:"title" gorm:"type:varchar(50);not null"`
	ContentText  string    `json:"question_text" gorm:"type:text;not null"`
	ContentImage string    `json:"question_image" gorm:"type:text"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;default:current_timestamp"`
}
