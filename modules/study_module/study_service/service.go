package study_service

import (
	"context"
	"fmt"
	"sync"

	"github.com/emel-study/emel-study-server/modules/study_module/study_models"
)

type SessionData struct {
	ID        string
	Positions []study_models.SoundPosition
	Answers   study_models.SessionAnswersRequest
}

type StudyService struct {
	mu       sync.RWMutex
	sessions map[string]*SessionData
	sounds   []study_models.SoundItem
}

func NewStudyService(sounds []study_models.SoundItem) *StudyService {
	return &StudyService{
		sessions: make(map[string]*SessionData),
		sounds:   sounds,
	}
}

func (s *StudyService) GetSounds(ctx context.Context) []study_models.SoundItem {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sounds
}

func (s *StudyService) CreateSession(ctx context.Context) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	n := len(s.sessions) + 1
	id := fmt.Sprintf("session-%d", n)
	for s.sessions[id] != nil {
		n++
		id = fmt.Sprintf("session-%d", n)
	}
	s.sessions[id] = &SessionData{ID: id}
	return id, nil
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
