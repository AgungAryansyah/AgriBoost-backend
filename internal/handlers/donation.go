package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/services"

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
		return err
	}

	var donation entity.Donation
	d.donationService.GetDonationById(&donation, param)
	return ctx.JSON(donation)
}

func (d *DonationHandler) GetDonationByUser(ctx *fiber.Ctx) error {
	var param dto.DonationParam

	if err := ctx.BodyParser(&param); err != nil {
		return err
	}

	var donation []entity.Donation
	d.donationService.GetDonationByUser(&donation, param)
	return ctx.JSON(donation)
}

func (d *DonationHandler) GetDonationByCampaign(ctx *fiber.Ctx) error {
	var param dto.DonationParam

	if err := ctx.BodyParser(&param); err != nil {
		return err
	}

	var donation []entity.Donation
	d.donationService.GetDonationByCampaign(&donation, param)
	return ctx.JSON(donation)
}
