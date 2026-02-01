package study_handler

import (
	"context"

	"github.com/emel-study/emel-study-server/modules/study_module/study_models"
	"github.com/gofiber/fiber/v2"
)

func (h *StudyHandlers) CreateSession(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID, err := h.service.CreateSession(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"sessionId": sessionID})
	}
}

func (h *StudyHandlers) SaveMap(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Params("id")
		if sessionID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing session id"})
		}
		var req study_models.SessionMapRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := h.service.SaveMap(ctx, sessionID, req.Positions); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"ok": true})
	}
}

func (h *StudyHandlers) SaveAnswers(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sessionID := c.Params("id")
		if sessionID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing session id"})
		}
		var req study_models.SessionAnswersRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if err := h.service.SaveAnswers(ctx, sessionID, req); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"ok": true})
	}
}
