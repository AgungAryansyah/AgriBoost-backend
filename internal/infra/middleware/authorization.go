package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AdminOnly(ctx *fiber.Ctx) error {
	authToken := ctx.GetReqHeaders()["Authorization"]

	if len(authToken) < 1 {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "token tidak valid",
		})
	}

	bearerToker := authToken[0]
	token := strings.Split(bearerToker, " ")

	id, isAdmin, err := m.jwt.ValidateToken(token[1])
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "token tidak valid",
		})
	}

	if !isAdmin {
		return ctx.Status(403).JSON(fiber.Map{
			"message": "tidak ada akses",
		})
	}

	ctx.Locals("userId", id)

	return ctx.Next()
}
