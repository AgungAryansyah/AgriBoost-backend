package dto

import "github.com/google/uuid"

type DonationParam struct {
	DonationId uuid.UUID `json:"donation_id" validated:"uuid"`
	CampaignId uuid.UUID `json:"campaign_id" validated:"uuid"`
	UserId     uuid.UUID `json:"user_id" validated:"uuid"`
}

type Donate struct {
	CampaignId uuid.UUID `json:"campaign_id" validated:"uuid"`
	UserId     uuid.UUID `json:"user_id" validated:"uuid"`
	Amount     int       `json:"amount" validated:"min=1000"`
}
