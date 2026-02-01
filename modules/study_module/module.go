package study_module

import (
	"context"

	"github.com/emel-study/emel-study-server/modules/study_module/study_constants"
	"github.com/emel-study/emel-study-server/modules/study_module/study_handler"
	"github.com/emel-study/emel-study-server/modules/study_module/study_models"
	"github.com/emel-study/emel-study-server/modules/study_module/study_service"
	"github.com/gofiber/fiber/v2"
)

type StudyModule struct {
	service  *study_service.StudyService
	handlers *study_handler.StudyHandlers
}

func NewStudyModule(ctx context.Context, app *fiber.App) (*StudyModule, error) {
	sounds := buildDefaultSounds()
	svc := study_service.NewStudyService(sounds)
	handlers, err := study_handler.NewStudyHandlers(app, svc)
	if err != nil {
		return nil, err
	}
	m := &StudyModule{service: svc, handlers: handlers}
	if err := m.handlers.Init(ctx); err != nil {
		return nil, err
	}
	return m, nil
}

func buildDefaultSounds() []study_models.SoundItem {
	ids := study_constants.DefaultSoundIDs()
	sounds := make([]study_models.SoundItem, len(ids))
	for i, id := range ids {
		// Placeholder: use same test tone or replace with real audio URLs from config
		sounds[i] = study_models.SoundItem{
			ID:       id,
			AudioURL: "/sounds/" + id + ".mp3", // client can proxy or use full URL
			Order:    i + 1,
		}
	}
	return sounds
}
