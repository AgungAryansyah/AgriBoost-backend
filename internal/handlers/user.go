package handlers

import (
	"AgriBoost/internal/infra/middleware"
	"AgriBoost/internal/models/dto"
	"AgriBoost/internal/services"
	"AgriBoost/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService services.UserServiceItf
	validator   *validator.Validate
	middleware  middleware.MiddlewareItf
}

func NewUserHandler(routerGroup fiber.Router, validator *validator.Validate, userService services.UserServiceItf, middleware middleware.MiddlewareItf) {
	UserHandler := UserHandler{
		userService: userService,
		validator:   validator,
		middleware:  middleware,
	}

	routerGroup = routerGroup.Group("/users")

	routerGroup.Post("/register", UserHandler.Register)
	routerGroup.Post("/login", UserHandler.Login)
	routerGroup.Patch("/edit", middleware.Authentication, UserHandler.EditProfile)
}

func (u *UserHandler) Register(ctx *fiber.Ctx) error {
	var register dto.Register
	if err := ctx.BodyParser(&register); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := u.validator.Struct(register); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	var id uuid.UUID
	if err := u.userService.Register(register, &id); err != nil {
		return utils.HttpError(ctx, "failed to create user", err)
	}

	return utils.HttpSuccess(ctx, "user created", id)
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	var login dto.Login
	if err := ctx.BodyParser(&login); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := u.validator.Struct(login); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	var id uuid.UUID
	token, err := u.userService.Login(login, &id)
	if err != nil {
		return utils.HttpError(ctx, "failed to log in", err)
	}

	return utils.HttpSuccess(ctx, "successfully login", token, id)
}

func (u *UserHandler) EditProfile(ctx *fiber.Ctx) error {
	var edit dto.EditProfile
	if err := ctx.BodyParser(&edit); err != nil {
		return utils.HttpError(ctx, "can't parse data, wrong JSON request format", err)
	}

	if err := u.validator.Struct(edit); err != nil {
		return utils.HttpError(ctx, "invalid data", err)
	}

	if err := u.userService.EditProfile(&edit); err != nil {
		return utils.HttpError(ctx, "failed to edit profile", err)
	}

	return utils.HttpSuccess(ctx, "success", nil)
}
