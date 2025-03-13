package dto

import "github.com/google/uuid"

type CreateCommunity struct {
	Name        string `json:"name" validate:"required,min=6"`
	Description string `json:"description" validate:"required,min=10"`
}

type JoinCommunity struct {
	CommunityID uuid.UUID `json:"community_id" validate:"uuid"`
	UserID      uuid.UUID `json:"user_id" validate:"uuid"`
}

type CommunityParam struct {
	CommunityID uuid.UUID `json:"community_id" validate:"uuid"`
	UserID      uuid.UUID `json:"user_id" validate:"uuid"`
	Name        string    `json:"name" validate:"min=6"`
}

type LeaveCommunity struct {
	CommunityID uuid.UUID `json:"community_id" validate:"uuid"`
	UserID      uuid.UUID `json:"user_id" validate:"uuid"`
}
