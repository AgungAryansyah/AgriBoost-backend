package repositories

import (
	entity "AgriBoost/internal/models/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageRepoItf interface {
	GetMessages(messages *[]entity.Message, page, pageSize int, communityId uuid.UUID) error
	CreateMessage(message *entity.Message) error
}

type MessageRepo struct {
	db *gorm.DB
}

func NewMessageRepo(db *gorm.DB) MessageRepoItf {
	return &MessageRepo{db}
}

func (m *MessageRepo) GetMessages(messages *[]entity.Message, page, pageSize int, communityId uuid.UUID) error {
	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	return m.db.Model(messages).
		Select("messages.message_id, messages.message, users.name AS username, messages.time_sent").
		Joins("JOIN users ON users.user_id = messages.user_id").
		Where("messages.community_id = ?", communityId).
		Order("messages.time_sent DESC").
		Limit(pageSize).Offset(offset).
		Scan(messages).Error
}

func (m *MessageRepo) CreateMessage(message *entity.Message) error {
	return m.db.Create(message).Error
}
