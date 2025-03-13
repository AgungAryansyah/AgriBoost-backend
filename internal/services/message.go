package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type MessageServiceItf interface {
	GetMessages(messages *[]dto.MessageDto, param dto.MessageParam) error
	SendMessage(msg string, communityId, userId uuid.UUID) error
}

type MessageService struct {
	messageRepo repositories.MessageRepoItf
}

func NewMessageService(messageRepo repositories.MessageRepoItf) MessageServiceItf {
	return &MessageService{
		messageRepo: messageRepo,
	}
}

func (m *MessageService) GetMessages(messages *[]dto.MessageDto, param dto.MessageParam) error {
	return m.messageRepo.GetMessages(messages, param.Page, param.PageSize, param.CommunityId)
}

func (m *MessageService) SendMessage(message string, communityId, userId uuid.UUID) error {
	return m.messageRepo.CreateMessage(&entity.Message{
		MessageID:   uuid.New(),
		Message:     message,
		CommunityId: communityId,
		UserId:      userId,
		TimeSent:    time.Now(),
	})
}
