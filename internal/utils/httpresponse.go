package utils

import (
	"github.com/gofiber/fiber/v2"
)

func HttpSuccess(c *fiber.Ctx, msg string, payload any) error {
	if payload == nil {
		return c.JSON(fiber.Map{
			"message": msg,
		})
	}

	return c.JSON(fiber.Map{
		"message": msg,
		"payload": payload,
	})
}

func HttpError(c *fiber.Ctx, msg string, err error) error {
	return c.Status(500).JSON(fiber.Map{
		"message": msg,
		"error":   err,
	})
}
