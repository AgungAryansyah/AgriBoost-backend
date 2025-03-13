package handlers

import (
	"AgriBoost/internal/models/dto"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService services.UserServiceItf
	validator   *validator.Validate
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, userService services.UserServiceItf) {
	UserHandler := UserHandler{
		userService: userService,
		validator:   validator,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", UserHandler.Register)
	routerGroup.Get("/login", UserHandler.Login)
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	var register dto.Register
	if err := ctx.BodyParser(&register); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := ctx.BodyParser(&register); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := u.userService.Register(register); err != nil {
		return utils.HttpError(ctx, "failed to create user", err)
	}

	return utils.HttpSuccess(ctx, "user created", nil)
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	var login dto.Login
	if err := ctx.BodyParser(&login); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := u.validator.Struct(login); err != nil {
		return utils.HttpError(ctx, "invalid request", err)
	}

	token, err := u.userService.Login(login)
	if err != nil {
		return utils.HttpError(ctx, "failed to log in", err)
	}

	return utils.HttpSuccess(ctx, "successfully login", token)
}
