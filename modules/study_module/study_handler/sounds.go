package study_handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func (h *StudyHandlers) GetSounds(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sounds := h.service.GetSounds(ctx)
		return c.JSON(fiber.Map{"sounds": sounds})
	}
}
