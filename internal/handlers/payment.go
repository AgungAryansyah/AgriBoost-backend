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

type PaymentHandler struct {
	paymentService services.PaymentServiceItf
	validator      *validator.Validate
	middleware     middleware.MiddlewareItf
}

func NewPaymentHandler(routerGroup fiber.Router, paymentService services.PaymentServiceItf, validator *validator.Validate, middleware middleware.MiddlewareItf) {
	PaymentHandler := PaymentHandler{
		paymentService: paymentService,
		validator:      validator,
		middleware:     middleware,
	}

	routerGroup = routerGroup.Group("/payment")
	routerGroup.Get("/id", middleware.AdminOnly, PaymentHandler.GetPaymentById)
	routerGroup.Get("/user", middleware.Authentication, PaymentHandler.GetpaymentByUser)
	routerGroup.Get("/all", middleware.AdminOnly, PaymentHandler.GetAllPayment)
}

func (p *PaymentHandler) GetPaymentById(ctx *fiber.Ctx) error {
	var param dto.PaymentParam

	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var payment entity.Payment
	err := p.paymentService.GetPaymentById(&payment, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", payment)
}

func (p *PaymentHandler) GetpaymentByUser(ctx *fiber.Ctx) error {
	var param dto.PaymentParam

	if err := ctx.BodyParser(&param); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	var payment []entity.Payment
	err := p.paymentService.GetPaymentByUser(&payment, param)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", payment)
}

func (p *PaymentHandler) GetAllPayment(ctx *fiber.Ctx) error {
	var payment []entity.Payment
	err := p.paymentService.GetAllPayment(&payment)

	if err != nil {
		return utils.HttpError(ctx, "failed to get data from the database", err)
	}

	return utils.HttpSuccess(ctx, "success", payment)
}
