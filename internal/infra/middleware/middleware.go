package middleware

import (
	"AgriBoost/internal/infra/jwt"

	"github.com/gofiber/fiber/v2"
)

type MiddlewareItf interface {
	Authentication(*fiber.Ctx) error
	AdminOnly(*fiber.Ctx) error
}

type Middleware struct {
	jwt jwt.JWTItf
}

func NewMiddleware(jwt jwt.JWTItf) MiddlewareItf {
	return &Middleware{
		jwt: jwt,
	}
}
