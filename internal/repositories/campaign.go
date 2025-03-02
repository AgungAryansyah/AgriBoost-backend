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

func (r *CampaignRepo) Get(campaign *[]entity.Campaign, campaignParam dto.CampaignParam) error {
	return r.db.Find(campaign, campaignParam).Error
}

func (r *CampaignRepo) GetOne(campaign *entity.Campaign, campaignParam dto.CampaignParam) error {
	return r.db.First(campaign, campaignParam).Error
}

func (r *CampaignRepo) Create(campaign *entity.Campaign) error {
	return r.db.Create(campaign).Error
}
