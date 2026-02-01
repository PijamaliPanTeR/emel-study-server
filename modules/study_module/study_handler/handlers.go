package study_handler

import (
	"context"

	"github.com/emel-study/emel-study-server/modules/study_module/study_service"
	"github.com/gofiber/fiber/v2"
)

type StudyHandlers struct {
	app     *fiber.App
	service *study_service.StudyService
}

func NewStudyHandlers(app *fiber.App, service *study_service.StudyService) (*StudyHandlers, error) {
	return &StudyHandlers{app: app, service: service}, nil
}

func (h *StudyHandlers) Init(ctx context.Context) error {
	h.app.Get("/study/sounds", h.GetSounds(ctx))
	h.app.Post("/study/session", h.CreateSession(ctx))
	h.app.Post("/study/session/:id/map", h.SaveMap(ctx))
	h.app.Post("/study/session/:id/answers", h.SaveAnswers(ctx))
	h.app.Post("/study/session/:id/progress", h.SaveProgress(ctx))
	return nil
}
