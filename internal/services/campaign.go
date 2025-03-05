package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"

	"github.com/google/uuid"
)

type CampaignServiceItf interface {
	GetCampaignById(campaignDto *dto.CampaignDto, campaignParam dto.CampaignParam) error
	GetCampaigns(campaignsDto *[]dto.CampaignDto, campaignParam dto.CampaignParam) error
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

func (c *CampaignService) GetCampaignById(campaignDto *dto.CampaignDto, campaignParam dto.CampaignParam) error {
	var campaign entity.Campaign
	if err := c.campaignRepo.GetOne(&campaign, campaignParam); err != nil {
		return err
	}

	dto.CampaignToDto(campaign, campaignDto)

	return nil
}

func (c *CampaignService) GetCampaigns(campaignsDto *[]dto.CampaignDto, campaignParam dto.CampaignParam) error {
	var campaigns []entity.Campaign
	if err := c.campaignRepo.Get(&campaigns, campaignParam); err != nil {
		return err
	}

	dto.CampaignsToDto(campaigns, campaignsDto)

	return nil
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
