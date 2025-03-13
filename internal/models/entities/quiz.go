package entity

import (
	"time"

	"github.com/google/uuid"
)

type Quiz struct {
	QuizId      uuid.UUID     `json:"quiz_id" gorm:"type:uuid;primaryKey"`
	Theme       string        `json:"theme" gorm:"type:varchar(50);not null"`
	Title       string        `json:"title" gorm:"type:varchar(50);not null"`
	Questions   []Question    `gorm:"foreignKey:QuizId"`
	QuizAttempt []QuizAttempt `gorm:"foreignKey:QuizId"`
}

type Question struct {
	QuestionId    uuid.UUID        `json:"question_id" gorm:"type:uuid;primaryKey"`
	QuizId        uuid.UUID        `json:"quiz_id" gorm:"type:uuid;not null"`
	Score         int              `json:"score" gorm:"type:integer;not null;check:score>0"`
	QuestionText  string           `json:"question_text" gorm:"type:text;not null"`
	QuestionImage string           `json:"question_image" gorm:"type:text"`
	Options       []QuestionOption `gorm:"foreignKey:QuestionId"`
}

type QuestionOption struct {
	OptionId    uuid.UUID `json:"option_id" gorm:"type:uuid;primaryKey"`
	QuestionId  uuid.UUID `json:"question_id" gorm:"type:uuid;not null"`
	IsCorrect   bool      `json:"is_correct" gorm:"type:bool;not null"`
	OptionText  string    `json:"option_text" gorm:"type:text;not null"`
	OptionImage string    `json:"option_image" gorm:"type:text"`
}

type QuizAttempt struct {
	AttemptId    uuid.UUID `json:"attempt_id" gorm:"type:uuid;primaryKey"`
	UserId       uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	QuizId       uuid.UUID `json:"quiz_id" gorm:"type:uuid;not null"`
	TotalScore   int       `json:"total_score" gorm:"type:integer;not null"`
	FinishedTime time.Time `json:"finished_time" gorm:"type:timestamp;default:current_timestamp"`
}
