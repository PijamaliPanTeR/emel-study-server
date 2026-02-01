package study_service

import (
	"context"

	"github.com/emel-study/emel-study-server/modules/study_module/study_models"
	"github.com/emel-study/emel-study-server/modules/study_module/study_repository"
)

type StudyService struct {
	repo study_repository.StudyRepository
}

func NewStudyService(repo study_repository.StudyRepository) *StudyService {
	return &StudyService{repo: repo}
}

func (s *StudyService) CreateSession(ctx context.Context, fingerprint string) (*study_models.CreateSessionResponse, error) {
	if fingerprint != "" {
		sess, err := s.repo.GetSessionByFingerprint(ctx, fingerprint)
		if err != nil {
			return nil, err
		}
		if sess != nil {
			resp := &study_models.CreateSessionResponse{SessionID: sess.ID, CurrentStep: sess.CurrentStep}
			if len(sess.Positions) > 0 {
				resp.Positions = sess.Positions
			}
			if sess.Answers.GroupStrategy != "" || sess.Answers.GroupsRepresent != "" {
				resp.Answers = &sess.Answers
			}
			if len(sess.ListenedSoundIDs) > 0 {
				resp.ListenedSoundIDs = sess.ListenedSoundIDs
			}
			if len(sess.SoundGroups) > 0 {
				resp.SoundGroups = sess.SoundGroups
			}
			if len(sess.DefineGroupsRectangles) > 0 {
				resp.DefineGroupsRectangles = sess.DefineGroupsRectangles
			}
			return resp, nil
		}
		_ = s.repo.DeleteFingerprint(ctx, fingerprint)
	}

	id, err := s.repo.NextSessionID(ctx)
	if err != nil {
		return nil, err
	}
	sess := &study_models.SessionData{ID: id, CurrentStep: "welcome"}
	if err := s.repo.UpsertSession(ctx, sess); err != nil {
		return nil, err
	}
	if fingerprint != "" {
		if err := s.repo.UpsertFingerprint(ctx, fingerprint, id); err != nil {
			return nil, err
		}
	}
	return &study_models.CreateSessionResponse{SessionID: id}, nil
}

func (s *StudyService) SaveMap(ctx context.Context, sessionID string, positions []study_models.SoundPosition) error {
	sess, err := s.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return err
	}
	if sess == nil {
		sess = &study_models.SessionData{ID: sessionID, CurrentStep: "welcome"}
	}
	sess.Positions = positions
	return s.repo.UpsertSession(ctx, sess)
}

func (s *StudyService) SaveAnswers(ctx context.Context, sessionID string, answers study_models.SessionAnswersRequest) error {
	sess, err := s.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return err
	}
	if sess == nil {
		sess = &study_models.SessionData{ID: sessionID, CurrentStep: "welcome"}
	}
	sess.Answers = answers
	return s.repo.UpsertSession(ctx, sess)
}

func (s *StudyService) SaveProgress(ctx context.Context, sessionID string, req study_models.SaveProgressRequest) error {
	sess, err := s.repo.GetSessionByID(ctx, sessionID)
	if err != nil {
		return err
	}
	if sess == nil {
		sess = &study_models.SessionData{ID: sessionID, CurrentStep: "welcome"}
	}
	if req.CurrentStep != "" {
		sess.CurrentStep = req.CurrentStep
	}
	if req.ListenedSoundIDs != nil {
		sess.ListenedSoundIDs = req.ListenedSoundIDs
	}
	if req.SoundGroups != nil {
		sess.SoundGroups = req.SoundGroups
	}
	if req.DefineGroupsRectangles != nil {
		sess.DefineGroupsRectangles = req.DefineGroupsRectangles
	}
	return s.repo.UpsertSession(ctx, sess)
}
