package study_module

import (
	"context"
	"database/sql"

	"github.com/emel-study/emel-study-server/modules/study_module/study_handler"
	"github.com/emel-study/emel-study-server/modules/study_module/study_repository"
	"github.com/emel-study/emel-study-server/modules/study_module/study_service"
	"github.com/gofiber/fiber/v2"
)

type StudyModule struct {
	service  *study_service.StudyService
	handlers *study_handler.StudyHandlers
}

func NewStudyModule(ctx context.Context, app *fiber.App, db *sql.DB) (*StudyModule, error) {
	repo, err := study_repository.NewStudyRepository(db)
	if err != nil {
		return nil, err
	}
	if err := repo.Init(ctx); err != nil {
		return nil, err
	}
	svc := study_service.NewStudyService(repo)
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
