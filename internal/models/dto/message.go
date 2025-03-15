package dto

import (
	"time"

	"github.com/google/uuid"
)

type MessageDto struct {
	Message  string    `json:"message"`
	UserName string    `json:"user_name"`
	TimeSent time.Time `json:"time_sent"`
}

type MessageParam struct {
	CommunityId uuid.UUID `json:"community_id" validate:"required,uuid"`
	Page        int       `json:"page" validate:"required,min=1"`
	PageSize    int       `json:"page_size" validate:"required,min=1"`
}

type SendMessage struct {
	Message     string    `validate:"required,string"`
	UserId      uuid.UUID `validate:"required,uuid"`
	CommunityId uuid.UUID `validate:"required,uuid"`
}
