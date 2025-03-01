package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/models/dto"
	entitiy "AgriBoost/internal/models/entities"
	"AgriBoost/internal/services"
	"net/http"

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
	var campaigns []entitiy.Campaign

	param := dto.CampaignParam{
		IsActive: true,
	}

	h.campaignService.GetCampaigns(&campaigns, param)
	return ctx.JSON(campaigns)
}

func (h *CampaignHandler) GetUserCampaign(ctx *fiber.Ctx) error {
	var campaigns []entitiy.Campaign
	var param dto.CampaignParam

	if err := ctx.BodyParser(&param); err != nil {
		return err
	}

	h.campaignService.GetCampaigns(&campaigns, param)
	return ctx.JSON(campaigns)
}

func (h *CampaignHandler) GetCampaign(ctx *fiber.Ctx) error {
	var param dto.CampaignParam

	if err := ctx.BodyParser(&param); err != nil {
		return err
	}

	var campaign entitiy.Campaign
	h.campaignService.GetCampaignById(&campaign, param)
	return ctx.JSON(campaign)
}

func (h *CampaignHandler) CreateCampaign(ctx *fiber.Ctx) error {
	var create dto.CreateCampaign

	if err := ctx.BodyParser(&create); err != nil {
		return err
	}

	if err := h.validator.Struct(create); err != nil {
		return err
	}

	err := h.campaignService.CreateCampaign(create)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK)
}
