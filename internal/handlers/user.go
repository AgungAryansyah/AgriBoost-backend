package handlers

import (
	"AgriBoost/internal/models/dto"
	"AgriBoost/internal/services"
	"net/http"

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

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	var register dto.Register

	if err := ctx.BodyParser(&register); err != nil {
		return err
	}

	if err := h.validator.Struct(register); err != nil {
		return err
	}

	err := h.userService.Register(register)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK)
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	var login dto.Login

	if err := ctx.BodyParser(&login); err != nil {
		return err
	}

	if err := h.validator.Struct(login); err != nil {
		return err
	}

	token, err := h.userService.Login(login)
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
