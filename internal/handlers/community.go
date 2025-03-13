package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CommunityHandler struct {
	communityService services.CommunityServiceItf
	validator        *validator.Validate
	middleware       middleware.MiddlewareItf
}

func NewCommunityHandler(routerGroup fiber.Router, validator *validator.Validate, communityService services.CommunityServiceItf, middleware middleware.MiddlewareItf) {
	communityHandler := CommunityHandler{
		communityService: communityService,
		validator:        validator,
		middleware:       middleware,
	}

	routerGroup = routerGroup.Group("community")
	routerGroup.Post("/create", middleware.AdminOnly, communityHandler.CreateCommunity)
	routerGroup.Get("/all", middleware.Authentication, communityHandler.GetAllCommunity)
	routerGroup.Get("/user", middleware.Authentication, communityHandler.GetUserCommunities)
	routerGroup.Post("/join", middleware.Authentication, communityHandler.JoinCommunity)
}

func (c *CommunityHandler) CreateCommunity(ctx *fiber.Ctx) error {
	var create dto.CreateCommunity
	if err := ctx.BodyParser(&create); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := c.validator.Struct(create); err != nil {
		return utils.HttpError(ctx, "invalid request", err)
	}

	if err := c.communityService.CreateCommunity(create); err != nil {
		return utils.HttpError(ctx, "failed to create community", err)
	}

	return utils.HttpSuccess(ctx, "community created", nil)
}

func (c *CommunityHandler) GetAllCommunity(ctx *fiber.Ctx) error {
	var communities []entity.Community
	if err := c.communityService.GetAllCommunity(&communities); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", communities)
}

func (c *CommunityHandler) GetUserCommunities(ctx *fiber.Ctx) error {
	var param dto.UserParam
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var communities []entity.Community
	if err := c.communityService.GetUserCommunities(&communities, param); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", communities)
}

func (c *CommunityHandler) JoinCommunity(ctx *fiber.Ctx) error {
	var param dto.JoinCommunity
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := c.communityService.JoinCommunity(param); err != nil {
		return utils.HttpError(ctx, "failed to join community", err)
	}

	return utils.HttpSuccess(ctx, "successfully joined a community", nil)
}

func (c *CommunityHandler) LeaveCommunity(ctx *fiber.Ctx) error {
	var leave dto.LeaveCommunity
	if err := ctx.BodyParser(&leave); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := c.communityService.LeaveCommunity(leave); err != nil {
		return utils.HttpError(ctx, "failed to join community", err)
	}

	return utils.HttpSuccess(ctx, "successfully leave a community", nil)
}
