package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommunityServiceItf interface {
	CreateCommunity(createCommunity dto.CreateCommunity) error
	GetAllCommunity(community *[]entity.Community) error
	GetUserCommunities(community *[]entity.Community, userParam dto.UserParam) error
	JoinCommunity(joinCommunity dto.JoinCommunity) error
	LeaveCommunity(leave dto.LeaveCommunity) error
	IsCommunityExist(exist *bool, comunityId uuid.UUID) error
}

type CommunityService struct {
	communityRepo repositories.CommunityRepoItf
	userRepo      repositories.UserRepoItf
}

func NewCommunityService(communityRepo repositories.CommunityRepoItf, userRepo repositories.UserRepoItf) CommunityServiceItf {
	return &CommunityService{
		communityRepo: communityRepo,
		userRepo:      userRepo,
	}
}

func (c *CommunityService) CreateCommunity(createCommunity dto.CreateCommunity) error {
	newCommunity := entity.Community{
		CommunityID: uuid.New(),
		Name:        createCommunity.Name,
		Description: createCommunity.Description,
	}

	return c.communityRepo.CreateCommunity(&newCommunity)
}

func (c *CommunityService) GetAllCommunity(community *[]entity.Community) error {
	return c.communityRepo.GetAllCommunity(community)
}

func (c *CommunityService) GetUserCommunities(community *[]entity.Community, userParam dto.UserParam) error {
	return c.communityRepo.GetUserCommunities(community, userParam)
}

func (c *CommunityService) JoinCommunity(joinCommunity dto.JoinCommunity) error {
	var user *entity.User
	if err := c.userRepo.IsUserExist(user, joinCommunity.UserID); err != nil {
		fmt.Println("error user dont exist")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user or community doesn't exist")
		}
		return err
	}

	var community *entity.Community
	if err := c.communityRepo.IsCommunityExist(community, joinCommunity.CommunityID); err != nil {
		fmt.Println("error community dont exist")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user or community doesn't exist")
		}
		return err
	}

	newCommunityMember := entity.CommunityMember{
		MemberID:    uuid.New(),
		UserId:      joinCommunity.UserID,
		CommunityId: joinCommunity.CommunityID,
	}

	return c.communityRepo.CreateCommunityMember(&newCommunityMember)
}

func (c *CommunityService) LeaveCommunity(leave dto.LeaveCommunity) error {
	var communityMember entity.CommunityMember
	if err := c.communityRepo.GetACommunityMember(&communityMember, leave.UserID, leave.CommunityID); err != nil {
		return err
	}

	return c.communityRepo.RemoveCommunityMember(&communityMember)
}

func (c *CommunityService) IsCommunityExist(exist *bool, comunityId uuid.UUID) error {
	var community *entity.Community
	if err := c.communityRepo.IsCommunityExist(community, comunityId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("community doesn't exist")
		}
		return err
	}

	*exist = true
	return nil
}
