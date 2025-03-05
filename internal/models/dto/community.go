package dto

import "github.com/google/uuid"

type CreateCommunity struct {
	Name        string `json:"name" validated:"required,min=6"`
	Description string `json:"description" validated:"required,min=10"`
}

type JoinCommunity struct {
	CommunityID uuid.UUID `json:"community_id" validated:"uuid"`
	UserID      uuid.UUID `json:"user_id" validated:"uuid"`
}

type CommunityParam struct {
	CommunityID uuid.UUID `json:"community_id" validated:"uuid"`
	Name        string    `json:"name" validated:"min=6"`
}
