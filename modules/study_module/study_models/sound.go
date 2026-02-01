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

// GroupEntry is one group: rectangle (bounds + soundIds) and generic per-group question answers (e.g. strategy, represent).
// Stored in group_info jsonb as an array of these.
type GroupEntry struct {
	Bounds   Bounds            `json:"bounds"`
	SoundIDs []string          `json:"soundIds"`
	Answers  map[string]string `json:"answers,omitempty"`
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
	SessionID        string          `json:"sessionId"`
	CurrentStep      string          `json:"currentStep,omitempty"`
	Positions        []SoundPosition `json:"positions,omitempty"`
	GroupInfo        []GroupEntry    `json:"groupInfo,omitempty"`
	ListenedSoundIDs []string        `json:"listenedSoundIds,omitempty"`
	SoundGroups      [][]string      `json:"soundGroups,omitempty"`
}

type SaveProgressRequest struct {
	CurrentStep      string       `json:"currentStep"`
	ListenedSoundIDs []string     `json:"listenedSoundIds,omitempty"`
	SoundGroups      [][]string   `json:"soundGroups,omitempty"`
	GroupInfo        []GroupEntry `json:"groupInfo,omitempty"`
}

type SessionData struct {
	ID               string
	Positions        []SoundPosition
	GroupInfo        []GroupEntry // single column: per-group bounds, soundIds, answers (strategy, represent, etc.)
	CurrentStep      string
	ListenedSoundIDs []string
	SoundGroups      [][]string
}
