package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/infra/midtrans"
	"AgriBoost/internal/models/dto"
	entity "AgriBoost/internal/models/entities"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DonationHandler struct {
	donationService services.DonationServiceItf
	validator       *validator.Validate
	middleware      middleware.MiddlewareItf
	midtrans        midtrans.MidtransItf
}

func NewDonationHandler(routerGroup fiber.Router, donationService services.DonationServiceItf, validator *validator.Validate, middleware middleware.MiddlewareItf, midtrans midtrans.MidtransItf) {
	DonationHandler := DonationHandler{
		donationService: donationService,
		validator:       validator,
		middleware:      middleware,
		midtrans:        midtrans,
	}

	routerGroup = routerGroup.Group("/donation")
	routerGroup.Post("/id", DonationHandler.GetDonationById)
	routerGroup.Post("/user", middleware.Authentication, DonationHandler.GetDonationByUser)
	routerGroup.Post("/campaign", DonationHandler.GetDonationByCampaign)
	routerGroup.Post("/donate", middleware.Authentication, DonationHandler.Donate)
	routerGroup.Post("/webhook", DonationHandler.HandleMidtransWebhook)
}

func (d *DonationHandler) GetDonationById(ctx *fiber.Ctx) error {
	var param dto.DonationParam
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := d.validator.Struct(param); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	var donation entity.Donation
	if err := d.donationService.GetDonationById(&donation, param); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", donation)
}

func (d *DonationHandler) GetDonationByUser(ctx *fiber.Ctx) error {
	var param dto.DonationParam
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := d.validator.Struct(param); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	var donation []entity.Donation
	if err := d.donationService.GetDonationByUser(&donation, param); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", donation)
}

func (d *DonationHandler) GetDonationByCampaign(ctx *fiber.Ctx) error {
	var param dto.DonationParam
	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := d.validator.Struct(param); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	var donation []entity.Donation
	if err := d.donationService.GetDonationByCampaign(&donation, param); err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", donation)
}

func (d *DonationHandler) Donate(ctx *fiber.Ctx) error {
	var donate dto.Donate
	if err := ctx.BodyParser(&donate); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := d.validator.Struct(donate); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	donationId := uuid.New()

	req, _ := d.midtrans.NewTransactionToken(donationId.String(), int64(donate.Amount))
	if req == nil {
		return utils.HttpError(ctx, "can't get transaction token", nil)
	}

	if err := d.donationService.Donate(donate, donationId); err != nil {
		return utils.HttpError(ctx, "can't store donation into the database", err)
	}

	return utils.HttpSuccess(ctx, "success", req)
}

func (d *DonationHandler) HandleMidtransWebhook(ctx *fiber.Ctx) error {
	var PaymentDetails map[string]interface{}
	if err := ctx.BodyParser(&PaymentDetails); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := d.donationService.HandleMidtransWebhook(PaymentDetails); err != nil {
		return utils.HttpError(ctx, "can't process payment details", err)
	}

	return utils.HttpSuccess(ctx, "success", nil)
}
