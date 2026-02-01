package server_handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type ServerHandlers struct {
	app *fiber.App
}

func NewServerHandlers(app *fiber.App) (*ServerHandlers, error) {
	return &ServerHandlers{app: app}, nil
}

func (h *ServerHandlers) Init(ctx context.Context) error {
	h.app.Get("/", h.HealthCheck(ctx))
	return nil
}
