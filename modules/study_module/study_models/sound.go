package study_models

type SoundItem struct {
	ID       string `json:"id"`
	Label    string `json:"label,omitempty"`
	AudioURL string `json:"audioUrl"`
	Order    int    `json:"order"`
}

type SoundPosition struct {
	SoundID string  `json:"soundId"`
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
}

type SessionMapRequest struct {
	Positions []SoundPosition `json:"positions"`
}

type SessionAnswersRequest struct {
	GroupStrategy   string `json:"groupStrategy"`
	GroupsRepresent string `json:"groupsRepresent"`
}

type CreateSessionRequest struct {
	Fingerprint string `json:"fingerprint"`
}

type DefineGroupRect struct {
	Bounds   Bounds   `json:"bounds"`
	SoundIDs []string `json:"soundIds"`
}

type Bounds struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type CreateSessionResponse struct {
	SessionID              string                 `json:"sessionId"`
	CurrentStep            string                 `json:"currentStep,omitempty"`
	Positions              []SoundPosition        `json:"positions,omitempty"`
	Answers                *SessionAnswersRequest `json:"answers,omitempty"`
	ListenedSoundIDs       []string               `json:"listenedSoundIds,omitempty"`
	SoundGroups            [][]string             `json:"soundGroups,omitempty"`
	DefineGroupsRectangles []DefineGroupRect      `json:"defineGroupsRectangles,omitempty"`
}

type SaveProgressRequest struct {
	CurrentStep            string            `json:"currentStep"`
	ListenedSoundIDs       []string          `json:"listenedSoundIds,omitempty"`
	SoundGroups            [][]string        `json:"soundGroups,omitempty"`
	DefineGroupsRectangles []DefineGroupRect `json:"defineGroupsRectangles,omitempty"`
}

type SessionData struct {
	ID                     string
	Positions              []SoundPosition
	Answers                SessionAnswersRequest
	CurrentStep            string
	ListenedSoundIDs       []string
	SoundGroups            [][]string
	DefineGroupsRectangles []DefineGroupRect
}
