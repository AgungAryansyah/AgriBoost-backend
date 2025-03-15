package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type MessageServiceItf interface {
	GetMessages(messages *[]entity.Message, param dto.MessageParam) error
	SendMessage(send dto.SendMessage) error
}

type MessageService struct {
	messageRepo repositories.MessageRepoItf
}

func NewMessageService(messageRepo repositories.MessageRepoItf) MessageServiceItf {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

func (m *MessageService) GetMessages(messages *[]entity.Message, param dto.MessageParam) error {
	return m.messageRepo.GetMessages(messages, param.Page, param.PageSize, param.CommunityId)
}

func (m *MessageService) SendMessage(send dto.SendMessage) error {
	return m.messageRepo.CreateMessage(&entity.Message{
		MessageID:   uuid.New(),
		Message:     send.Message,
		CommunityId: send.CommunityId,
		UserId:      send.UserId,
		TimeSent:    time.Now(),
	})
}
