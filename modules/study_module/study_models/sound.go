package study_models

type SoundItem struct {
	ID      string `json:"id"`
	Label   string `json:"label,omitempty"`
	AudioURL string `json:"audioUrl"`
	Order   int    `json:"order"`
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
