package server_handler

import (
	"context"
	"fmt"

	"github.com/emel-study/emel-study-server/pkg"
	"github.com/gofiber/fiber/v2"
)

func (h *ServerHandlers) HealthCheck(ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		_ = pkg.GetCurrentFuncName()
		return c.JSON(fiber.Map{
			"msg": fmt.Sprintf("Server is running with version: %s", pkg.Version),
		})
	}
}
