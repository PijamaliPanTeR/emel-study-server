package study_service

import (
	"context"
	"fmt"
	"sync"

	"github.com/emel-study/emel-study-server/modules/study_module/study_models"
)

type SessionData struct {
	ID                     string
	Positions              []study_models.SoundPosition
	Answers                study_models.SessionAnswersRequest
	CurrentStep            string
	ListenedSoundIDs       []string
	SoundGroups            [][]string
	DefineGroupsRectangles []study_models.DefineGroupRect
}

type StudyService struct {
	mu                sync.RWMutex
	sessions          map[string]*SessionData
	userByFingerprint map[string]string // fingerprint -> sessionID
	sounds            []study_models.SoundItem
}

func NewStudyService(sounds []study_models.SoundItem) *StudyService {
	return &StudyService{
		sessions:          make(map[string]*SessionData),
		userByFingerprint: make(map[string]string),
		sounds:            sounds,
	}
}

func (s *StudyService) GetSounds(ctx context.Context) []study_models.SoundItem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sounds
}

func (s *StudyService) CreateSession(ctx context.Context, fingerprint string) (*study_models.CreateSessionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if fingerprint != "" {
		if sid, ok := s.userByFingerprint[fingerprint]; ok {
			if sess, ok := s.sessions[sid]; ok {
				resp := &study_models.CreateSessionResponse{SessionID: sid, CurrentStep: sess.CurrentStep}
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
			delete(s.userByFingerprint, fingerprint)
		}
	}

	n := len(s.sessions) + 1
	id := fmt.Sprintf("session-%d", n)
	for s.sessions[id] != nil {
		n++
		id = fmt.Sprintf("session-%d", n)
	}
	s.sessions[id] = &SessionData{ID: id, CurrentStep: "welcome"}
	if fingerprint != "" {
		s.userByFingerprint[fingerprint] = id
	}
	return &study_models.CreateSessionResponse{SessionID: id}, nil
}

func (s *StudyService) SaveMap(ctx context.Context, sessionID string, positions []study_models.SoundPosition) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[sessionID]
	if !ok {
		sess = &SessionData{ID: sessionID}
		s.sessions[sessionID] = sess
	}
	sess.Positions = positions
	return nil
}

func (s *StudyService) SaveAnswers(ctx context.Context, sessionID string, answers study_models.SessionAnswersRequest) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[sessionID]
	if !ok {
		sess = &SessionData{ID: sessionID}
		s.sessions[sessionID] = sess
	}
	sess.Answers = answers
	return nil
}

func (s *StudyService) SaveProgress(ctx context.Context, sessionID string, req study_models.SaveProgressRequest) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[sessionID]
	if !ok {
		sess = &SessionData{ID: sessionID}
		s.sessions[sessionID] = sess
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
	return nil
}
