package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"

	"gorm.io/gorm"
)

type CampaignRepoItf interface {
	GetOne(campaign *entity.Campaign, campaignParam dto.CampaignParam) error
	Get(campaign *[]entity.Campaign, campaignParam dto.CampaignParam) error
	Create(campaign *entity.Campaign) error
}

type CampaignRepo struct {
	db *gorm.DB
}

func NewCampaignRepo(db *gorm.DB) CampaignRepoItf {
	return &CampaignRepo{db}
}

func (c *CampaignRepo) Get(campaign *[]entity.Campaign, campaignParam dto.CampaignParam) error {
	return c.db.Find(campaign, campaignParam).Error
}

func (c *CampaignRepo) GetOne(campaign *entity.Campaign, campaignParam dto.CampaignParam) error {
	return c.db.First(campaign, campaignParam).Error
}

func (c *CampaignRepo) Create(campaign *entity.Campaign) error {
	return c.db.Create(campaign).Error
}
