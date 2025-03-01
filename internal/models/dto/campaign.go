package dto

import "github.com/google/uuid"

type CampaignParam struct {
	CampaignId uuid.UUID `json:"campaign_id" validated:"uuid"`
	IsActive   bool
	UserId     uuid.UUID `json:"user_id" validated:"uuid"`
}

type CreateCampaign struct {
	Title       string    `json:"title" validated:"required,min=3"`
	Description string    `json:"description" validated:"required,min=10"`
	GoalAmount  int       `json:"goal_amount" validated:"required,min=1000"`
	UserId      uuid.UUID `json:"user_id" validated:"required,uuid"`
}
