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
	CommunityId uuid.UUID `json:"community_id" validated:"required,uuid"`
	Page        int       `json:"page" validated:"required,int"`
	PageSize    int       `json:"page_size" validated:"required,int"`
}

type SendMessage struct {
	Message     string    `validated:"required,string"`
	UserId      uuid.UUID `validated:"required,uuid"`
	CommunityId uuid.UUID `validated:"required,uuid"`
}
