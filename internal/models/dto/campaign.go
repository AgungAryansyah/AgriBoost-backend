package dto

import (
	entity "AgriBoost/internal/models/entities"
	"sync"
	"time"

	"github.com/google/uuid"
)

type CampaignParam struct {
	CampaignId uuid.UUID `json:"campaign_id" validate:"uuid"`
	IsActive   bool
	UserId     uuid.UUID `json:"user_id" validate:"uuid"`
}

type CreateCampaign struct {
	Title       string    `json:"title" validate:"required,min=3"`
	Description string    `json:"description" validate:"required,min=10"`
	GoalAmount  int       `json:"goal_amount" validate:"required,min=1000"`
	UserId      uuid.UUID `json:"user_id" validate:"required,uuid"`
}

type CampaignDto struct {
	CampaignId      uuid.UUID `json:"campaign_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	GoalAmount      int       `json:"goal_amount"`
	CollectedAmount int       `json:"collected_amount"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
	UserId          uuid.UUID `json:"user_id"`
}

func CampaignToDto(campaign entity.Campaign, campaignDto *CampaignDto) {
	*campaignDto = CampaignDto{
		CampaignId:      campaign.CampaignId,
		Title:           campaign.Title,
		Description:     campaign.Description,
		GoalAmount:      campaign.GoalAmount,
		CollectedAmount: campaign.CollectedAmount,
		IsActive:        campaign.IsActive,
		CreatedAt:       campaign.CreatedAt,
		UserId:          campaign.UserId,
	}
}

func CampaignsToDto(campaigns []entity.Campaign, campaignsDto *[]CampaignDto) {
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	*campaignsDto = make([]CampaignDto, len(campaigns))

	for i, val := range campaigns {
		wg.Add(1)
		go func(i int, val entity.Campaign) {
			defer wg.Done()

			dto := CampaignDto{}
			CampaignToDto(val, &dto)

			mu.Lock()
			(*campaignsDto)[i] = dto
			mu.Unlock()
		}(i, val)
	}

	wg.Wait()
}
