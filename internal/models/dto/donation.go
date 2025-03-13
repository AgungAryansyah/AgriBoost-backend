package dto

import "github.com/google/uuid"

type DonationParam struct {
	DonationId uuid.UUID `json:"donation_id" validate:"uuid"`
	CampaignId uuid.UUID `json:"campaign_id" validate:"uuid"`
	UserId     uuid.UUID `json:"user_id" validate:"uuid"`
}

type Donate struct {
	CampaignId uuid.UUID `json:"campaign_id" validate:"required,uuid"`
	UserId     uuid.UUID `json:"user_id" validate:"required,uuid"`
	Amount     int       `json:"amount" validate:"required,min=1000"`
}
