package study_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/emel-study/emel-study-server/modules/study_module/study_models"
)

func (r *StudyRepositoryImpl) GetSessionByID(ctx context.Context, sessionID string) (*study_models.SessionData, error) {
	q := `SELECT id, current_step, positions, group_info, listened_sound_ids, sound_groups
		FROM study_sessions WHERE id = $1`
	var (
		id              string
		currentStep     string
		positionsJSON   sql.NullString
		groupInfoJSON   sql.NullString
		listenedJSON    sql.NullString
		soundGroupsJSON sql.NullString
	)
	err := r.DB.QueryRowContext(ctx, q, sessionID).Scan(
		&id, &currentStep, &positionsJSON, &groupInfoJSON, &listenedJSON, &soundGroupsJSON,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	sess := &study_models.SessionData{ID: id, CurrentStep: currentStep}
	if positionsJSON.Valid && positionsJSON.String != "" {
		_ = json.Unmarshal([]byte(positionsJSON.String), &sess.Positions)
	}
	if groupInfoJSON.Valid && groupInfoJSON.String != "" {
		_ = json.Unmarshal([]byte(groupInfoJSON.String), &sess.GroupInfo)
	}
	if listenedJSON.Valid && listenedJSON.String != "" {
		_ = json.Unmarshal([]byte(listenedJSON.String), &sess.ListenedSoundIDs)
	}
	if soundGroupsJSON.Valid && soundGroupsJSON.String != "" {
		_ = json.Unmarshal([]byte(soundGroupsJSON.String), &sess.SoundGroups)
	}
	return sess, nil
}

func (r *StudyRepositoryImpl) GetSessionByFingerprint(ctx context.Context, fingerprint string) (*study_models.SessionData, error) {
	var sessionID string
	err := r.DB.QueryRowContext(ctx, `SELECT session_id FROM study_fingerprints WHERE fingerprint = $1`, fingerprint).Scan(&sessionID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.GetSessionByID(ctx, sessionID)
}

func (r *StudyRepositoryImpl) UpsertSession(ctx context.Context, sess *study_models.SessionData) error {
	positionsJSON, _ := json.Marshal(sess.Positions)
	groupInfoJSON, _ := json.Marshal(sess.GroupInfo)
	listenedJSON, _ := json.Marshal(sess.ListenedSoundIDs)
	soundGroupsJSON, _ := json.Marshal(sess.SoundGroups)

	q := `INSERT INTO study_sessions (id, current_step, positions, group_info, listened_sound_ids, sound_groups, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, now())
		ON CONFLICT (id) DO UPDATE SET
			current_step = EXCLUDED.current_step,
			positions = EXCLUDED.positions,
			group_info = EXCLUDED.group_info,
			listened_sound_ids = EXCLUDED.listened_sound_ids,
			sound_groups = EXCLUDED.sound_groups,
			updated_at = now()`
	_, err := r.DB.ExecContext(ctx, q,
		sess.ID, sess.CurrentStep, string(positionsJSON), string(groupInfoJSON),
		string(listenedJSON), string(soundGroupsJSON),
	)
	return err
}

func (r *StudyRepositoryImpl) UpsertFingerprint(ctx context.Context, fingerprint, sessionID string) error {
	_, err := r.DB.ExecContext(ctx,
		`INSERT INTO study_fingerprints (fingerprint, session_id) VALUES ($1, $2)
		 ON CONFLICT (fingerprint) DO UPDATE SET session_id = EXCLUDED.session_id`,
		fingerprint, sessionID,
	)
	return err
}

func (r *StudyRepositoryImpl) DeleteFingerprint(ctx context.Context, fingerprint string) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM study_fingerprints WHERE fingerprint = $1`, fingerprint)
	return err
}

func (r *StudyRepositoryImpl) NextSessionID(ctx context.Context) (string, error) {
	var next int64
	err := r.DB.QueryRowContext(ctx,
		`SELECT COALESCE(MAX(CAST(REGEXP_REPLACE(id, '^session-', '') AS BIGINT)), 0) + 1 FROM study_sessions WHERE id ~ '^session-[0-9]+$'`).Scan(&next)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("session-%d", next), nil
}
