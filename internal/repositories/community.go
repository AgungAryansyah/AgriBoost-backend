package repositories

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommunityRepoItf interface {
	CreateCommunity(community *entity.Community) error
	GetAllCommunity(community *[]entity.Community) error
	GetUserCommunities(community *[]entity.Community, userParam dto.UserParam) error
	CreateCommunityMember(communityMember *entity.CommunityMember) error
	IsCommunityExist(exist *bool, communityId uuid.UUID) error
}

type CommunityRepo struct {
	db *gorm.DB
}

func NewCommunityRepo(db *gorm.DB) CommunityRepoItf {
	return &CommunityRepo{db}
}

func (c *CommunityRepo) CreateCommunity(community *entity.Community) error {
	return c.db.Create(community).Error
}

func (c *CommunityRepo) GetAllCommunity(community *[]entity.Community) error {
	return c.db.Find(community).Error
}

func (c *CommunityRepo) GetUserCommunities(community *[]entity.Community, userParam dto.UserParam) error {
	return c.db.Find(community, userParam).Error
}

func (c *CommunityRepo) CreateCommunityMember(communityMember *entity.CommunityMember) error {
	return c.db.Create(communityMember).Error
}

func (c *CommunityRepo) IsCommunityExist(exist *bool, communityId uuid.UUID) error {
	return c.db.Model(&entity.Community{}).Select("community_id").Where("community_id = ?", communityId).First(&exist).Error
}
