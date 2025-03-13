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
	GetUserCommunities(community *[]entity.Community, communityParam dto.CommunityParam) error
	CreateCommunityMember(communityMember *entity.CommunityMember) error
	IsCommunityExist(community *entity.Community, communityId uuid.UUID) error
	RemoveCommunityMember(member *entity.CommunityMember) error
	GetACommunityMember(member *entity.CommunityMember, userId uuid.UUID, communityId uuid.UUID) error
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

func (c *CommunityRepo) GetUserCommunities(communities *[]entity.Community, communityParam dto.CommunityParam) error {
	return c.db.Joins("JOIN community_members cm ON cm.community_id = communities.community_id").
		Where("cm.user_id = ?", communityParam.UserID).
		Find(&communities).Error
}

func (c *CommunityRepo) CreateCommunityMember(communityMember *entity.CommunityMember) error {
	return c.db.Create(communityMember).Error
}

func (c *CommunityRepo) IsCommunityExist(community *entity.Community, communityId uuid.UUID) error {
	return c.db.Model(&entity.Community{}).Select("community_id").Where("community_id = ?", communityId).First(&community).Error
}

func (c *CommunityRepo) RemoveCommunityMember(member *entity.CommunityMember) error {
	return c.db.Delete(member).Error
}

func (c *CommunityRepo) GetACommunityMember(member *entity.CommunityMember, userId uuid.UUID, communityId uuid.UUID) error {
	return c.db.Where("user_id = ? AND community_id = ?", userId, communityId).First(member).Error
}
