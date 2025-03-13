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
	CommunityId uuid.UUID `json:"community_id"`
	Page        int       `json:"page"`
	PageSize    int       `json:"page_size"`
}

type SendMessage struct {
	Message     string    `json:"message"`
	UserId      uuid.UUID `json:"user_id"`
	CommunityId uuid.UUID `json:"community_id"`
}
