package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"

	"github.com/google/uuid"
)

type CampaignServiceItf interface {
	GetCampaignById(campaign *entity.Campaign, campaignParam dto.CampaignParam) error
	GetCampaigns(campaigns *[]entity.Campaign, campaignParam dto.CampaignParam) error
	CreateCampaign(create dto.CreateCampaign) error
}

type CampaignService struct {
	campaignRepo repositories.CampaignRepoItf
}

func NewCampaignService(campaignRepo repositories.CampaignRepoItf) CampaignServiceItf {
	return &CampaignService{
		campaignRepo: campaignRepo,
	}
}

func (c *CampaignService) GetCampaignById(campaign *entity.Campaign, campaignParam dto.CampaignParam) error {
	return c.campaignRepo.GetOne(campaign, campaignParam)
}

func (c *CampaignService) GetCampaigns(campaigns *[]entity.Campaign, campaignParam dto.CampaignParam) error {
	return c.campaignRepo.Get(campaigns, campaignParam)
}
func (c *CampaignService) CreateCampaign(create dto.CreateCampaign) error {
	NewCampaign := entity.Campaign{
		CampaignId:  uuid.New(),
		Title:       create.Title,
		Description: create.Description,
		GoalAmount:  create.GoalAmount,
		UserId:      create.UserId,
	}
	return c.campaignRepo.Create(&NewCampaign)
}
