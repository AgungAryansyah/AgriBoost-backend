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
