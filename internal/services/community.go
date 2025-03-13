package services

import (
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/repositories"
	"errors"

	"github.com/google/uuid"
)

type CommunityServiceItf interface {
	CreateCommunity(createCommunity dto.CreateCommunity) error
	GetAllCommunity(community *[]entity.Community) error
	GetUserCommunities(community *[]entity.Community, userParam dto.UserParam) error
	JoinCommunity(joinCommunity dto.JoinCommunity) error
	LeaveCommunity(leave dto.LeaveCommunity) error
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
	var userExist bool
	if err := c.userRepo.IsUserExist(&userExist, joinCommunity.UserID); err != nil {
		return err
	}

	var communityExist bool
	if err := c.communityRepo.IsCommunityExist(&communityExist, joinCommunity.CommunityID); err != nil {
		return err
	}

	if !userExist || !communityExist {
		return errors.New("user or community doesn't exist")
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
