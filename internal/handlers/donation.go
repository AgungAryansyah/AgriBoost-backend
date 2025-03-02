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

type DonationHandler struct {
	donationService services.DonationServiceItf
	validator       *validator.Validate
	middleware      middleware.MiddlewareItf
}

func NewDonationgHandler(routerGroup fiber.Router, donationServie services.DonationServiceItf, validator *validator.Validate, middleware middleware.MiddlewareItf) {
	DonationHandler := DonationHandler{
		donationService: donationServie,
		validator:       validator,
		middleware:      middleware,
	}

	routerGroup = routerGroup.Group("/donation")
	routerGroup.Get("/id", DonationHandler.GetDonationById)
	routerGroup.Get("/user", middleware.Authentication, DonationHandler.GetDonationByUser)
	routerGroup.Get("/campaign", DonationHandler.GetDonationByCampaign)
}

func (d *DonationHandler) GetDonationById(ctx *fiber.Ctx) error {
	var param dto.DonationParam

	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var donation entity.Donation
	err := d.donationService.GetDonationById(&donation, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", donation)
}

func (d *DonationHandler) GetDonationByUser(ctx *fiber.Ctx) error {
	var param dto.DonationParam

	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var donation []entity.Donation
	err := d.donationService.GetDonationByUser(&donation, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", donation)
}

func (d *DonationHandler) GetDonationByCampaign(ctx *fiber.Ctx) error {
	var param dto.DonationParam

	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var donation []entity.Donation
	err := d.donationService.GetDonationByCampaign(&donation, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", donation)
}
