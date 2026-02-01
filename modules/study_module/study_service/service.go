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
			if len(sess.GroupInfo) > 0 {
				resp.GroupInfo = sess.GroupInfo
			}
			if len(sess.ListenedSoundIDs) > 0 {
				resp.ListenedSoundIDs = sess.ListenedSoundIDs
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
	// Merge into first group's answers so legacy client doesn't lose data
	if len(sess.GroupInfo) > 0 {
		if sess.GroupInfo[0].Answers == nil {
			sess.GroupInfo[0].Answers = make(map[string]string)
		}
		if answers.GroupStrategy != "" {
			sess.GroupInfo[0].Answers["strategy"] = answers.GroupStrategy
		}
		if answers.GroupsRepresent != "" {
			sess.GroupInfo[0].Answers["represent"] = answers.GroupsRepresent
		}
	} else {
		sess.GroupInfo = []study_models.GroupEntry{{
			Answers: map[string]string{
				"strategy":  answers.GroupStrategy,
				"represent": answers.GroupsRepresent,
			},
		}}
	}
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
	if req.GroupInfo != nil {
		sess.GroupInfo = req.GroupInfo
	}
	return s.repo.UpsertSession(ctx, sess)
}
