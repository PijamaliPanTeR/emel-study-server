package server_module

import (
	"context"

	"github.com/emel-study/emel-study-server/modules/server_module/server_handler"
	"github.com/gofiber/fiber/v2"
)

type ServerModule struct {
	handlers *server_handler.ServerHandlers
}

func NewServerModule(ctx context.Context, app *fiber.App) (*ServerModule, error) {
	handlers, err := server_handler.NewServerHandlers(app)
	if err != nil {
		return nil, err
	}
	m := &ServerModule{handlers: handlers}
	if err := m.handlers.Init(ctx); err != nil {
		return nil, err
	}
	return m, nil
}
