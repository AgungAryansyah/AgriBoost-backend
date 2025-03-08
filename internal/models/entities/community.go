package entity

import (
	"github.com/google/uuid"
)

type Community struct {
	CommunityID     uuid.UUID         `json:"community_id" gorm:"type:uuid;primaryKey"`
	Name            string            `json:"name" gorm:"type:varchar(255);not null;unique"`
	Description     string            `json:"description" gorm:"type:text;not null"`
	CommunityMember []CommunityMember `gorm:"foreignKey:CommunityId"`
}

type CommunityMember struct {
	MemberID    uuid.UUID `json:"member_id" gorm:"type:uuid;primaryKey"`
	UserId      uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CommunityId uuid.UUID `json:"community_id" gorm:"type:uuid;not null"`
}
