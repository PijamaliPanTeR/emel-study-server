package study_repository

import (
	"context"
	"database/sql"

	"github.com/emel-study/emel-study-server/modules/study_module/study_models"
)

// StudyRepository defines persistence for study sessions and fingerprints (same structure as main-service repositories).
type StudyRepository interface {
	Init(ctx context.Context) error
	CreateStudySessionsTable(ctx context.Context, db *sql.DB) error
	CreateStudyFingerprintsTable(ctx context.Context, db *sql.DB) error
	GetSessionByID(ctx context.Context, sessionID string) (*study_models.SessionData, error)
	GetSessionByFingerprint(ctx context.Context, fingerprint string) (*study_models.SessionData, error)
	UpsertSession(ctx context.Context, sess *study_models.SessionData) error
	UpsertFingerprint(ctx context.Context, fingerprint, sessionID string) error
	DeleteFingerprint(ctx context.Context, fingerprint string) error
	NextSessionID(ctx context.Context) (string, error)
}

type StudyRepositoryImpl struct {
	DB *sql.DB
}

func NewStudyRepository(db *sql.DB) (StudyRepository, error) {
	return &StudyRepositoryImpl{DB: db}, nil
}

func (r *StudyRepositoryImpl) Init(ctx context.Context) error {
	if err := r.CreateStudySessionsTable(ctx, r.DB); err != nil {
		return err
	}
	return r.CreateStudyFingerprintsTable(ctx, r.DB)
}

func (r *StudyRepositoryImpl) CreateStudySessionsTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, CreateStudySessionsTableQuery)
	return err
}

func (r *StudyRepositoryImpl) CreateStudyFingerprintsTable(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, CreateStudyFingerprintsTableQuery); err != nil {
		return err
	}
	_, err := db.ExecContext(ctx, CreateStudyFingerprintsIndexQuery)
	return err
}
