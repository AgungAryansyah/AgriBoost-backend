package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/models/dto"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CampaignHandler struct {
	campaignService services.CampaignServiceItf
	validator       *validator.Validate
	middleware      middleware.MiddlewareItf
}

func NewCampaignHandler(routerGroup fiber.Router, validator *validator.Validate, campaignService services.CampaignServiceItf, middleware middleware.MiddlewareItf) {
	CampaignHandler := CampaignHandler{
		campaignService: campaignService,
		validator:       validator,
		middleware:      middleware,
	}

	routerGroup = routerGroup.Group("/campaign")

	routerGroup.Get("/active", CampaignHandler.GetActiveCampaign)
	routerGroup.Get("/get", CampaignHandler.GetCampaign)
	routerGroup.Get("/user", middleware.Authentication, CampaignHandler.GetUserCampaign)
	routerGroup.Post("/create", middleware.Authentication, CampaignHandler.CreateCampaign)
}

func (h *CampaignHandler) GetActiveCampaign(ctx *fiber.Ctx) error {
	var campaignsDto []dto.CampaignDto

	param := dto.CampaignParam{
		IsActive: true,
	}

	err := h.campaignService.GetCampaigns(&campaignsDto, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", campaignsDto)
}

func (h *CampaignHandler) GetUserCampaign(ctx *fiber.Ctx) error {
	var param dto.CampaignParam

	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var campaignsDto []dto.CampaignDto
	err := h.campaignService.GetCampaigns(&campaignsDto, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", campaignsDto)
}

func (h *CampaignHandler) GetCampaign(ctx *fiber.Ctx) error {
	var param dto.CampaignParam

	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var campaignDto dto.CampaignDto
	err := h.campaignService.GetCampaignById(&campaignDto, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", campaignDto)
}

func (h *CampaignHandler) CreateCampaign(ctx *fiber.Ctx) error {
	var create dto.CreateCampaign

	if err := ctx.BodyParser(&create); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := h.validator.Struct(create); err != nil {
		return utils.HttpError(ctx, "invalid request", err)
	}

	err := h.campaignService.CreateCampaign(create)
	if err != nil {
		return utils.HttpError(ctx, "failed to create campaign", err)
	}

	return utils.HttpSuccess(ctx, "campaign created", nil)
}
